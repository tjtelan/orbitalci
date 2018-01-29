package werker

import (
	d "bitbucket.org/level11consulting/go-til/deserialize"
	ocelog "bitbucket.org/level11consulting/go-til/log"
	pb "bitbucket.org/level11consulting/ocelot/protos"
	"bitbucket.org/level11consulting/ocelot/util/storage"
	"bitbucket.org/level11consulting/ocelot/util/storage/models"
	b "bitbucket.org/level11consulting/ocelot/werker/builder"
	"github.com/golang/protobuf/proto"
	"time"
)

// Transport struct is for the Transport channel that will interact with the streaming side of the service
// to stream results back to the admin. It sends just enough to be unique, the hash that triggered the build
// and the InfoChan which the builder will write to.
type Transport struct {
	Hash     string
	InfoChan chan []byte
	DbId     int64
}

type WorkerMsgHandler struct {
	Topic        string
	WerkConf     *WerkerConf
	infochan     chan []byte
	ChanChan     chan *Transport
	Deserializer d.Deserializer
	Basher	     *b.Basher
	Store        storage.OcelotStorage
}

// UnmarshalAndProcess is called by the nsq consumer to handle the build message
func (w WorkerMsgHandler) UnmarshalAndProcess(msg []byte) error {
	ocelog.Log().Debug("unmarshaling build obj and processing")
	werkerTask := &pb.WerkerTask{}
	if err := proto.Unmarshal(msg, werkerTask); err != nil {
		ocelog.IncludeErrField(err).Warning("unmarshal error")
		return err
	}
	// channels get closed after the build finishes
	w.infochan = make(chan []byte)
	// set goroutine for watching for results and logging them (for rn)
	// cant add go watchForResults here bc can't call method on interface until it's been cast properly.

	var builder b.Builder
	switch w.WerkConf.werkerType {
	case Docker:
		builder = b.NewDockerBuilder(w.Basher)
	default:
		builder = b.NewDockerBuilder(w.Basher)
	}

	go w.MakeItSo(werkerTask, builder)
	return nil
}

// watchForResults sends the *Transport object over the transport channel for stream functions to process
func (w *WorkerMsgHandler) WatchForResults(hash string, dbId int64) {
	ocelog.Log().Debugf("adding hash ( %s ) & infochan to transport channel", hash)
	transport := &Transport{Hash: hash, InfoChan: w.infochan, DbId: dbId}
	w.ChanChan <- transport
}

// MakeItSo will call appropriate builder functions
func (w *WorkerMsgHandler) MakeItSo(werk *pb.WerkerTask, builder b.Builder) {
	ocelog.Log().Debug("hash build ", werk.CheckoutHash)
	//defers are stacked, will be executed FILO

	defer close(w.infochan)
	defer builder.Cleanup()
	//TODO: write stages to db
	//TODO: write build data to db

	w.WatchForResults(werk.CheckoutHash, werk.Id)
	var stageResults []*b.Result

	setupResult := builder.Setup(w.infochan, werk)
	stageResults = append(stageResults, setupResult)
	// todo: this is causing panic: runtime error: invalid memory address or nil pointer dereference
	// on the builder.Cleanup
	if setupResult.Status == b.FAIL {
		ocelog.Log().Error(setupResult.Error)
		return
	}
	fail := false
	start := time.Now()
	for _, stage := range werk.BuildConf.Stages {
		stageResult := builder.Execute(stage, w.infochan, werk.CheckoutHash)
		ocelog.Log().Info("finished stage", stage.Name)
		stageResults = append(stageResults, stageResult)
		// todo: should this check go before or after special build stuffs? i guess if it fails, there won't be
		// any deployment
		if stageResult.Status == b.FAIL {
			fail = true
			ocelog.Log().Error(stageResult.Error)
			// todo: update failure reasons
			if err := w.Store.AddFail(convertResultToFailureReasons(stageResult, werk.Id)); err != nil {
				ocelog.IncludeErrField(err).Error("couldn't add failure reasons")
			}
			break
		}
		// build is special because we deploy after this
		if stage.Name == "build" {
			// todo: deploy to nexus
			continue
		}
	}
	dura := time.Now().Sub(start)
	if err := w.Store.UpdateSum(fail, dura.Seconds(), werk.Id); err != nil {
		ocelog.IncludeErrField(err).Error("couldn't update summary in database")
	}

	ocelog.Log().Debugf("finished building id %s", werk.CheckoutHash)
}


func convertResultToFailureReasons(res *b.Result, id int64) *models.BuildFailureReason {
	var err string
	if res.Error == nil {
		err = ""
	} else {
		err = res.Error.Error()
	}
	fr :=  &models.FailureReasons{
				Stage: res.Stage,
				Status: int32(res.Status),
				Error: err,
				Messages: res.Messages,
			}
	return &models.BuildFailureReason{
		BuildId: id,
		FailureReasons: fr,
	}
}