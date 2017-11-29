package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/shankj3/ocelot/util/consulet"
	"github.com/shankj3/ocelot/util/ocelog"
	"github.com/shankj3/ocelot/util/ocenet"
	"github.com/shankj3/ocelot/util/storage"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{}
	consulDonePath = "ci/builds/%s/done" //  %s is hash
)
// modified from https://elithrar.github.io/article/custom-handlers-avoiding-globals/
type appContext struct {
	conf 		  *WerkerConf
	storage 	  storage.BuildOutputStorage
	buildInfo     map[string]*buildDatum
	consul        *consulet.Consulet
}


type buildDatum struct {
	buildData [][]byte
	done      chan int
}


func (a *appContext)  SetBuildDone(gitHash string) error {
	// todo: add byte of 0/1.. have ot use binary library though and idk how to use that yet
	// and not motivated enough to do it right now
	err := a.consul.AddKeyValue(fmt.Sprintf(consulDonePath, gitHash), []byte("true"))
	if err != nil {
		return err
	}
	return nil
}

func (a *appContext) CheckIfBuildDone(gitHash string) bool {
	kv, err := a.consul.GetKeyValue(fmt.Sprintf(consulDonePath, gitHash))
	if err != nil {
		// idk what we should be doing if the error is not nil, maybe panic? hope that never happens?
		return false
	}
	if kv != nil {
		return true
	}
	return false
}


func stream(ctx interface{}, w http.ResponseWriter, r *http.Request){
	a := ctx.(*appContext)
	vars := mux.Vars(r)
	hash := vars["hash"]
	ocelog.Log().Debug(hash)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ocelog.IncludeErrField(err).Error("wtf?")
		return
	}
	defer ws.Close()
	pumpDone := make(chan int)

	go pumpBundle(ws, a, hash, pumpDone)
	ocelog.Log().Debug("sending infoChan over web socket, waiting for the channel to be closed.")
	<-pumpDone
}

// kinda hackily done because switching to grpc
type WebsocketEy interface {
	SetWriteDeadline(t time.Time) error
	WriteMessage(messageType int, data []byte) error
	Close() error
}

// what to return to socket if something went awry
func writeWSError(ws WebsocketEy, description []byte) {
	ws.SetWriteDeadline(time.Now().Add(10*time.Second))
	ws.WriteMessage(websocket.TextMessage, []byte("ERROR!\n"))
	ws.WriteMessage(websocket.TextMessage, description)
	ws.Close()
}

// pumpBundle writes build data to web socket
func pumpBundle(ws WebsocketEy, appCtx *appContext, hash string, done chan int){
	// determine whether to get from storage or off infoReader
	if appCtx.CheckIfBuildDone(hash) {
		ocelog.Log().Debugf("build %s is done, getting from appCtx", hash)
		err := pumpFromStorage(appCtx, hash, ws)
		if err != nil {
			ocelog.IncludeErrField(err).Error("error retrieving from storage")
		}
	} else {
		ocelog.Log().Debug("pumping info array data to web socket")
		buildInfo, ok := appCtx.buildInfo[hash]
		if ok {
			err := streamFromArray(buildInfo, ws)
			if err != nil {
				ocelog.IncludeErrField(err).Error("could not stream from array!")
			}
			ocelog.Log().Debug("streamed build data from array")
		} else {
			writeWSError(ws, []byte("did not find hash in current streaming data and the build was not marked as done"))
		}
	}
	defer func(){
		close(done)
		ws.SetWriteDeadline(time.Now().Add(10 * time.Second))
		ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		ws.Close()
	}()
}

func pumpFromStorage(appCtx *appContext, hash string, ws WebsocketEy) error {
	reader, err := appCtx.storage.RetrieveReader(hash)
	if err != nil {
		ocelog.IncludeErrField(err).Warn("could not retrieve persisted build data")
		writeWSError(ws, []byte("could not retrieve persisted build data"))
		return err
	}
	s := bufio.NewScanner(reader)
	// write to web socket
	for s.Scan() {
		ws.SetWriteDeadline(time.Now().Add(10*time.Second))
		if err := ws.WriteMessage(websocket.TextMessage, s.Bytes()); err != nil{
			ocelog.IncludeErrField(err).Error("could not write to web socket")
			ws.Close()
			break
		}
	}
	return s.Err()
}

func streamFromArray(buildInfo *buildDatum, ws WebsocketEy) (err error){
	var index int
	var previousIndex int
	for {
		select {
		case <-buildInfo.done:
			return nil
		default:
			buildData := buildInfo.buildData[index:]
			ind, err := iterateOverBuildData(buildData, ws)
			previousIndex = index
			index += ind + 1
			ocelog.Log().WithField("lines_sent", ind).WithField("index", index).WithField("previousIndex", previousIndex).Debug()
			//fmt.Println("------------------------------------------------------------------")
			//fmt.Println("byte arrays sent   :  ", ind)
			//fmt.Println("index is at        :  ", index)
			//fmt.Println("previousIndex is at:  ", previousIndex)
			//fmt.Println("------------------------------------------------------------------")
			//time.Sleep(4*time.Second)
			if err != nil {
				return err
			}
		}
		time.Sleep(1 * time.Second)
	}

}


func iterateOverBuildData(data [][]byte, ws WebsocketEy) (int, error) {
	var index int
	for ind, dataLine := range data {
		ws.SetWriteDeadline(time.Now().Add(10*time.Second))
		if err := ws.WriteMessage(websocket.TextMessage, dataLine); err != nil {
			ocelog.IncludeErrField(err).Error("could not write to web socket")
			ws.Close()
			return index, err
		}
		index = ind
	}
	return index, nil
}

// writeInfoChanToInMemMap is what processes transport objects that come from the transport channel (objects get created
// 	and sent when a new build is pulled off the queue).
// 	the info channel is written to an array which is put in a map in the appCtx along with a done channel so
//  there is a way to see when the array will not be written to anymore
//  when the info channel is closed and the loop finishes, all the data is written to the storage defined in the
//  appCtx, the done flag is written to consul, and the array is removed from the map
func writeInfoChanToInMemMap(transport  *Transport, appCtx *appContext){
	var dataSlice [][]byte
	build := &buildDatum{dataSlice, make(chan int),}
	appCtx.buildInfo[transport.Hash] = build
	ocelog.Log().Debugf("writing infochan data for %s", transport.Hash)
	// todo: change the worker to act as grpc server
	// to expose method that lets you get stream as commit
	// keep in memory of output of build
	// create a new stream @ every request
	for i := range transport.InfoChan {
		build.buildData = append(build.buildData, i)
	}
	ocelog.Log().Debug("supposedly done???")
	build.done <- 0
	err := appCtx.storage.StoreLines(transport.Hash, build.buildData)
	if err != nil {
		ocelog.IncludeErrField(err).Error("could not store build data to storage")
		return
	}
	// get rid of hash from cache, set build done in consul
	if err := appCtx.SetBuildDone(transport.Hash); err != nil {
		ocelog.IncludeErrField(err).Error("could not set build done")
	}
	ocelog.Log().Debugf("removing hash %s from readerCache and channelDict", transport.Hash)
	delete(appCtx.buildInfo, transport.Hash)

}

func cacheProcessor(transpo chan *Transport, appCtx *appContext){
	for i := range transpo {
		ocelog.Log().Debugf("adding info channel for hash %s to map for streaming access.", i.Hash)
		go writeInfoChanToInMemMap(i, appCtx)
	}
}


func serveHome(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "test.html")
}

func GetWerkConfig(conf *WerkerConf) *appContext{
	store := storage.NewFileBuildStorage("")
	appCtx := &appContext{ conf: conf, storage: store, buildInfo: make(map[string]*buildDatum), consul: consulet.Default()}
	return appCtx
}

func ServeMe(transportChan chan *Transport, conf *WerkerConf){
	ocelog.Log().Debug("started serving routine for streaming data")
	appctx := GetWerkConfig(conf)
	go cacheProcessor(transportChan, appctx)
	muxi := mux.NewRouter()
	muxi.Handle("/ws/builds/{hash}", &ocenet.AppContextHandler{appctx, stream}).Methods("GET")
	muxi.HandleFunc("/builds/{hash}", serveHome).Methods("GET")
	n := ocenet.InitNegroni("werker", muxi)
	n.Run(":"+conf.servicePort)
	ocelog.Log().Info("QUITTING")
}