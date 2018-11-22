// coconut_worker.go - Worker for Coconut client.
// Copyright (C) 2018  Jedrzej Stuczynski.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package coconutclient implements operations required by a client of coconut IA server.
package coconutclient

import (
	"github.com/jstuczyn/CoconutGo/logger"
	"gopkg.in/op/go-logging.v1"

	"fmt"
	"sync"

	"github.com/jstuczyn/CoconutGo/crypto/coconut/concurrency/jobpacket"
	"github.com/jstuczyn/CoconutGo/crypto/coconut/scheme"
	"github.com/jstuczyn/CoconutGo/server/commands"
	"github.com/jstuczyn/CoconutGo/worker"
)

// Worker allows writing coconut actions to a shared job queue,
// so that they could be run concurrently.
type Worker struct {
	worker.Worker

	incomingCh <-chan interface{}
	jobQueue   chan<- interface{}
	log        *logging.Logger

	muxParams *MuxParams
	sk        *coconut.SecretKey // ensure they can be safely shared between multiple workers
	vk        *coconut.VerificationKey

	avk *coconut.VerificationKey // only used if server is a provider

	id uint64
}

// AddToJobQueue adds a job packet directly to the job queue.
// currently for testing sake; todo: should I use this instead of writing manually?
func (ccw *Worker) AddToJobQueue(jobpacket *jobpacket.JobPacket) {
	ccw.jobQueue <- jobpacket
}

func (ccw *Worker) worker() {
	for {
		var cmdReq *commands.CommandRequest
		select {
		case <-ccw.HaltCh():
			ccw.log.Debugf("Halting Coconut worker %d\n", ccw.id)
			return
		case e := <-ccw.incomingCh:
			cmdReq = e.(*commands.CommandRequest)
			cmd := cmdReq.Cmd()
			var respData interface{}
			respStatus := commands.StatusCode_UNKNOWN
			errMsg := ""

			switch v := cmd.(type) {
			case *commands.Sign:
				ccw.log.Notice("Received Sign (NOT blind) command")
				if len(v.PubM()) > len(ccw.sk.Y()) {
					errMsg = fmt.Sprintf("Received more attributes to sign than what the server supports. Got: %v, expected at most: %v", len(v.PubM()), len(ccw.sk.Y()))
					ccw.log.Error(errMsg)
					respData = nil
					respStatus = commands.StatusCode_INVALID_ARGUMENTS
					continue
				}
				sig, err := ccw.Sign(ccw.muxParams, ccw.sk, v.PubM())
				if err != nil {
					// todo: should client really know those details?
					errMsg = fmt.Sprintf("Error while signing message: %v", err)
					ccw.log.Errorf(errMsg)
					respData = nil
					respStatus = commands.StatusCode_PROCESSING_ERROR
					continue
				}
				ccw.log.Debugf("Writing back signature")
				respData = sig
			case *commands.Vk:
				ccw.log.Notice("Received Get Verification Key command")
				respData = ccw.vk
			case *commands.Verify:
				ccw.log.Notice("Received Verify (NOT blind) command")
				if ccw.avk != nil {
					respData = ccw.Verify(ccw.muxParams, ccw.avk, v.PubM(), v.Sig())
				} else {
					errMsg = "The aggregate verification key is nil. Is the server a provider? And if so, has it completed the start up sequence?"
					ccw.log.Error(errMsg)
					respData = nil
					respStatus = commands.StatusCode_UNAVAILABLE
				}
			case *commands.BlindSign:
				ccw.log.Notice("Received Blind Sign command")
				if len(v.PubM())+len(v.BlindSignMats().Enc()) > len(ccw.sk.Y()) {
					errMsg = fmt.Sprintf("Received more attributes to sign than what the server supports. Got: %v, expected at most: %v", len(v.PubM())+len(v.BlindSignMats().Enc()), len(ccw.sk.Y()))
					ccw.log.Error(errMsg)
					respData = nil
					respStatus = commands.StatusCode_INVALID_ARGUMENTS
					continue
				}
				sig, err := ccw.BlindSign(ccw.muxParams, ccw.sk, v.BlindSignMats(), v.EgPub(), v.PubM())
				if err != nil {
					// todo: should client really know those details?
					errMsg = fmt.Sprintf("Error while signing message: %v", err)
					ccw.log.Errorf(errMsg)
					respData = nil
					respStatus = commands.StatusCode_PROCESSING_ERROR
					continue
				}
				ccw.log.Debugf("Writing back blinded signature")
				respData = sig
			case *commands.BlindVerify:
				ccw.log.Notice("Received Blind Verify Command")
				if ccw.avk != nil {
					respData = ccw.BlindVerify(ccw.muxParams, ccw.avk, v.Sig(), v.BlindShowMats(), v.PubM())
				} else {
					errMsg = "The aggregate verification key is nil. Is the server a provider? And if so, has it completed the start up sequence?"
					ccw.log.Error(errMsg)
					respData = nil
					respStatus = commands.StatusCode_UNAVAILABLE
				}
			default:
				errMsg = "Received Invalid Command"
				ccw.log.Critical(errMsg)
				respStatus = commands.StatusCode_INVALID_COMMAND
			}

			cmdReq.RetCh() <- &commands.Response{
				Data:         respData,
				ErrorStatus:  respStatus,
				ErrorMessage: errMsg,
			}
		}
	}
}

// New creates new instance of a coconutWorker.
// todo: simplify attributes...
// nolint: lll
func New(jobQueue chan<- interface{}, incomingCh <-chan interface{}, id uint64, l *logger.Logger, params *coconut.Params, sk *coconut.SecretKey, vk *coconut.VerificationKey, avk *coconut.VerificationKey) *Worker {
	// params are passed rather than generated by the clientworker, as each client would waste cpu cycles by generating
	// the same values + they HAD TO be pregenerated anyway in order to create the keys
	muxParams := &MuxParams{params, sync.Mutex{}}
	w := &Worker{
		jobQueue:   jobQueue,
		incomingCh: incomingCh,
		id:         id,
		muxParams:  muxParams,
		sk:         sk,
		vk:         vk,
		avk:        avk,
		log:        l.GetLogger(fmt.Sprintf("CoconutClientWorker:%d", int(id))),
	}

	w.Go(w.worker)
	return w
}

// func init with q to make params
