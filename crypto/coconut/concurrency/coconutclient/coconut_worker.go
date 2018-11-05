// todo: A LOT of refractoring and moving stuff around

package coconutclient

import (
	"github.com/jstuczyn/CoconutGo/logger"
	"gopkg.in/op/go-logging.v1"

	"fmt"
	"sync"

	"github.com/jstuczyn/CoconutGo/crypto/coconut/concurrency/jobpacket"
	coconut "github.com/jstuczyn/CoconutGo/crypto/coconut/scheme"
	"github.com/jstuczyn/CoconutGo/server/commands"
	"github.com/jstuczyn/CoconutGo/worker"
)

// Worker allows writing coconut actions to a shared job queue,
// so that they could be run concurrently.
// todo: introduce more attributes as needed, perhaps keep params here?
type Worker struct {
	worker.Worker
	params   *MuxParams
	initOnce sync.Once

	incomingCh <-chan interface{}
	jobQueue   chan<- interface{}
	log        *logging.Logger

	muxParams *MuxParams
	sk        *coconut.SecretKey // ensure they can be safely shared between multiple workers
	vk        *coconut.VerificationKey

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
			switch v := cmd.(type) {
			case *commands.Sign:
				ccw.log.Debug("Sign cmd")
				if len(v.PubM()) > len(ccw.sk.Y()) {
					ccw.log.Error("Too many params to sign.")
					cmdReq.RetCh() <- nil
					continue
				}
				sig, err := ccw.Sign(ccw.muxParams, ccw.sk, v.PubM())
				if err != nil {
					ccw.log.Error("Error while signing message")
					cmdReq.RetCh() <- err
					continue
				}
				ccw.log.Debugf("Writing back signature")
				cmdReq.RetCh() <- sig
			case *commands.Vk:
				ccw.log.Debug("Get Vk cmd")
				cmdReq.RetCh() <- ccw.vk
			}
		}
	}
}

// New creates new instance of a coconutWorker.
// todo: simplify attributes...
func New(jobQueue chan<- interface{}, incomingCh <-chan interface{}, id uint64, l *logger.Logger, params *coconut.Params, sk *coconut.SecretKey, vk *coconut.VerificationKey) *Worker {
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
		log:        l.GetLogger(fmt.Sprintf("CoconutClientWorker:%d", int(id))),
	}

	w.Go(w.worker)
	return w
}

// func init with q to make params
