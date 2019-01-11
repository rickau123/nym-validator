// cryptoworker.go - Coconut Worker for Coconut server.
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

// Package cryptoworker gives additional functionalities to regular CoconutWorker
// that are required by a server instance.
package cryptoworker

import (
	"fmt"

	"0xacab.org/jstuczyn/CoconutGo/common/comm/commands"
	"0xacab.org/jstuczyn/CoconutGo/crypto/coconut/concurrency/coconutworker"
	"0xacab.org/jstuczyn/CoconutGo/crypto/coconut/concurrency/jobpacket"
	"0xacab.org/jstuczyn/CoconutGo/crypto/coconut/scheme"
	"0xacab.org/jstuczyn/CoconutGo/crypto/elgamal"
	"0xacab.org/jstuczyn/CoconutGo/logger"
	"0xacab.org/jstuczyn/CoconutGo/worker"
	"gopkg.in/op/go-logging.v1"
)

const (
	defaultErrorMessage    = ""
	defaultErrorStatusCode = commands.StatusCode_UNKNOWN

	providerStartupErr = "The aggregate verification key is nil. " +
		"Is the server a provider? And if so, has it completed the start up sequence?"
)

// CryptoWorker allows writing coconut actions to a shared job queue,
// so that they could be run concurrently.
type CryptoWorker struct {
	worker.Worker
	*coconutworker.CoconutWorker // TODO: since coconutWorker is created in New, does it need to be a reference?

	incomingCh <-chan *commands.CommandRequest
	log        *logging.Logger

	sk *coconut.SecretKey // ensure they can be safely shared between multiple workers
	vk *coconut.VerificationKey

	avk *coconut.VerificationKey // only used if server is a provider

	id uint64
}

func getDefaultResponse() *commands.Response {
	return &commands.Response{
		Data:         nil,
		ErrorStatus:  defaultErrorStatusCode,
		ErrorMessage: defaultErrorMessage,
	}
}

func (cw *CryptoWorker) setErrorResponse(response *commands.Response, errMsg string, errCode commands.StatusCode) {
	cw.log.Error(errMsg)
	response.Data = nil
	response.ErrorMessage = errMsg
	response.ErrorStatus = errCode
}

func (cw *CryptoWorker) handleSignRequest(req *commands.SignRequest) *commands.Response {
	response := getDefaultResponse()

	if len(req.PubM) > len(cw.sk.Y()) {
		errMsg := fmt.Sprintf("Received more attributes to sign than what the server supports."+
			" Got: %v, expected at most: %v", len(req.PubM), len(cw.sk.Y()))
		cw.setErrorResponse(response, errMsg, commands.StatusCode_INVALID_ARGUMENTS)
		return response
	}
	sig, err := cw.SignWrapper(cw.sk, coconut.BigSliceFromByteSlices(req.PubM))
	if err != nil {
		// TODO: should client really know those details?
		errMsg := fmt.Sprintf("Error while signing message: %v", err)
		cw.setErrorResponse(response, errMsg, commands.StatusCode_PROCESSING_ERROR)
		return response
	}
	cw.log.Debugf("Writing back signature")
	response.Data = sig
	return response
}

func (cw *CryptoWorker) handleVerificationKeyRequest(req *commands.VerificationKeyRequest) *commands.Response {
	response := getDefaultResponse()
	response.Data = cw.vk
	return response
}

func (cw *CryptoWorker) handleVerifyRequest(req *commands.VerifyRequest) *commands.Response {
	response := getDefaultResponse()

	if cw.avk == nil {
		errMsg := providerStartupErr
		cw.setErrorResponse(response, errMsg, commands.StatusCode_UNAVAILABLE)
		return response
	}
	sig := &coconut.Signature{}
	if err := sig.FromProto(req.Sig); err != nil {
		errMsg := "Could not recover received signature."
		cw.setErrorResponse(response, errMsg, commands.StatusCode_INVALID_ARGUMENTS)
		return response
	}
	response.Data = cw.VerifyWrapper(cw.avk, coconut.BigSliceFromByteSlices(req.PubM), sig)
	return response
}

func (cw *CryptoWorker) handleBlindSignRequest(req *commands.BlindSignRequest) *commands.Response {
	response := getDefaultResponse()

	lambda := &coconut.Lambda{}
	if err := lambda.FromProto(req.Lambda); err != nil {
		errMsg := "Could not recover received lambda."
		cw.setErrorResponse(response, errMsg, commands.StatusCode_INVALID_ARGUMENTS)
		return response
	}
	if len(req.PubM)+len(lambda.Enc()) > len(cw.sk.Y()) {
		errMsg := fmt.Sprintf("Received more attributes to sign than what the server supports."+
			" Got: %v, expected at most: %v", len(req.PubM)+len(lambda.Enc()), len(cw.sk.Y()))
		cw.setErrorResponse(response, errMsg, commands.StatusCode_INVALID_ARGUMENTS)
		return response
	}
	egPub := &elgamal.PublicKey{}
	if err := egPub.FromProto(req.EgPub); err != nil {
		errMsg := "Could not recover received ElGamal Public Key."
		cw.setErrorResponse(response, errMsg, commands.StatusCode_INVALID_ARGUMENTS)
		return response
	}
	sig, err := cw.BlindSignWrapper(cw.sk, lambda, egPub, coconut.BigSliceFromByteSlices(req.PubM))
	if err != nil {
		// TODO: should client really know those details?
		errMsg := fmt.Sprintf("Error while signing message: %v", err)
		cw.setErrorResponse(response, errMsg, commands.StatusCode_PROCESSING_ERROR)
		return response
	}
	cw.log.Debugf("Writing back blinded signature")
	response.Data = sig
	return response
}

func (cw *CryptoWorker) handleBlindVerifyRequest(req *commands.BlindVerifyRequest) *commands.Response {
	response := getDefaultResponse()

	if cw.avk == nil {
		errMsg := providerStartupErr
		cw.setErrorResponse(response, errMsg, commands.StatusCode_UNAVAILABLE)
		return response
	}
	sig := &coconut.Signature{}
	if err := sig.FromProto(req.Sig); err != nil {
		errMsg := "Could not recover received signature."
		cw.setErrorResponse(response, errMsg, commands.StatusCode_INVALID_ARGUMENTS)
		return response
	}
	theta := &coconut.Theta{}
	if err := theta.FromProto(req.Theta); err != nil {
		errMsg := "Could not recover received theta."
		cw.setErrorResponse(response, errMsg, commands.StatusCode_INVALID_ARGUMENTS)
		return response
	}
	response.Data = cw.BlindVerifyWrapper(cw.avk, sig, theta, coconut.BigSliceFromByteSlices(req.PubM))
	return response
}

// nolint: gocyclo
func (cw *CryptoWorker) worker() {
	for {
		select {
		case <-cw.HaltCh():
			cw.log.Noticef("Halting Coconut Server worker %d\n", cw.id)
			return
		case cmdReq := <-cw.incomingCh:
			cmd := cmdReq.Cmd()
			var response *commands.Response

			switch req := cmd.(type) {
			case *commands.SignRequest:
				cw.log.Notice("Received Sign (NOT blind) command")
				response = cw.handleSignRequest(req)

			case *commands.VerificationKeyRequest:
				cw.log.Notice("Received Get Verification Key command")
				response = cw.handleVerificationKeyRequest(req)

			case *commands.VerifyRequest:
				cw.log.Notice("Received Verify (NOT blind) command")
				response = cw.handleVerifyRequest(req)

			case *commands.BlindSignRequest:
				cw.log.Notice("Received Blind Sign command")
				response = cw.handleBlindSignRequest(req)

			case *commands.BlindVerifyRequest:
				cw.log.Notice("Received Blind Verify Command")
				response = cw.handleBlindVerifyRequest(req)

			default:
				errMsg := "Received Invalid Command"
				cw.log.Critical(errMsg)
				response = getDefaultResponse()
				response.ErrorStatus = commands.StatusCode_INVALID_COMMAND
			}
			cmdReq.RetCh() <- response
		}
	}
}

// Config encapsulates arguments passed in New to create new instance of the cryptoworker.
type Config struct {
	JobQueue   chan<- *jobpacket.JobPacket
	IncomingCh <-chan *commands.CommandRequest

	ID uint64

	Log *logger.Logger

	Params *coconut.Params
	Sk     *coconut.SecretKey
	Vk     *coconut.VerificationKey
	Avk    *coconut.VerificationKey
}

// New creates new instance of a coconutWorker.
func New(cfg *Config) *CryptoWorker {
	cw := &CryptoWorker{
		CoconutWorker: coconutworker.New(cfg.JobQueue, cfg.Params),
		incomingCh:    cfg.IncomingCh,
		id:            cfg.ID,
		sk:            cfg.Sk,
		vk:            cfg.Vk,
		avk:           cfg.Avk,
		log:           cfg.Log.GetLogger(fmt.Sprintf("Servercryptoworker:%d", int(cfg.ID))),
	}

	cw.Go(cw.worker)
	return cw
}
