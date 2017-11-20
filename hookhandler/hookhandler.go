package main

import (
	"github.com/gorilla/mux"
	"github.com/meatballhat/negroni-logrus"
	"github.com/shankj3/ocelot/admin/handler"
	"github.com/shankj3/ocelot/admin/models"
	pb "github.com/shankj3/ocelot/protos/out"

	//"github.com/shankj3/ocelot/util/consulet"
	"github.com/shankj3/ocelot/util/deserialize"
	"github.com/shankj3/ocelot/util/nsqpb"
	"github.com/shankj3/ocelot/util/ocelog"
	"github.com/shankj3/ocelot/util/ocenet"
	"github.com/shankj3/ocelot/util/ocevault"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
	"os"
	"sync"
)

//var consul = consulet.Default()
var deserializer = deserialize.New()

// vault singleton
var vaultCached *ocevault.Ocevault
var vaultOnce sync.Once

// producer singleton
var producerCached *nsqpb.PbProduce
var producerOnce sync.Once


// On receive of repo push, marshal the json to an object then write the important fields to protobuf Message on NSQ queue.
func RepoPush(w http.ResponseWriter, r *http.Request) {
	repopush := &pb.RepoPush{}
	if err := deserializer.JSONToProto(r.Body, repopush); err != nil {
		ocenet.JSONApiError(w, http.StatusBadRequest, "could not parse request body into proto.Message", err)
	}
	fullName := repopush.Repository.FullName
	hash := repopush.Push.Changes[0].New.Target.Hash
	buildConf, err := GetBuildConfig(fullName, hash)
	if err != nil {
		// if the build file just isn't there don't worry about it.
		if err != ocenet.FileNotFound {
			ocenet.JSONApiError(w, http.StatusBadRequest,"unable to get build conf", err)
			return
		}
		ocelog.Log().Debugf("no ocelot yml found for repo %s", repopush.Repository.FullName)
		return
	}
	ocelog.Log().Debug("found ocelot yml")
	vault := ocevault.GetInitVault(vaultOnce, vaultCached)
	token, err := vault.CreateThrowawayToken()
	if err != nil {
		ocenet.JSONApiError(w, http.StatusBadRequest, "unable to create one-time vault token", err)
	}
	// instead, add to topic. each worker gets a topic off a channel,
	// so one worker to one channel
	bundle := &pb.PushBuildBundle{
		Config:     buildConf,
		PushData:   repopush,
		VaultToken: token,
	}
	ocelog.Log().Debug("created push bundle")
	pbProducer := nsqpb.GetInitProducer(producerOnce, producerCached)
	go pbProducer.WriteToNsq(bundle, nsqpb.PushTopic)
}

func PullRequest(w http.ResponseWriter, r *http.Request) {
	pr := &pb.PullRequest{}
	if err := deserializer.JSONToProto(r.Body, pr); err != nil {
		ocenet.JSONApiError(w, http.StatusBadRequest, "could not parse request body into proto.Message", err)
		return
	}
	fullName := pr.Pullrequest.Source.Repository.FullName
	hash := pr.Pullrequest.Source.Commit.Hash

	buildConf, err := GetBuildConfig(fullName, hash)
	if err != nil {
		//ocelog.IncludeErrField(err).Error("unable to get build conf")
		// if the build file just isn't there don't worry about it.
		if err != ocenet.FileNotFound {
			ocenet.JSONApiError(w, http.StatusBadRequest, "unable to get build conf", err)
			return
		}
		ocelog.Log().Debugf("no ocelot yml found for repo %s", pr.Pullrequest.Source.Repository.FullName)
		return
	}
	// get one-time token use for access to vault
	vault := ocevault.GetInitVault(vaultOnce, vaultCached)
	token, err := vault.CreateThrowawayToken()
	if err != nil {
		ocenet.JSONApiError(w, http.StatusBadRequest, "unable to create one-time vault token", err)
		return
	}
	// create bundle, send that s*** off!
	bundle := &pb.PRBuildBundle{
		Config:     buildConf,
		PrData:     pr,
		VaultToken: token,
	}
	pbProducer := nsqpb.GetInitProducer(producerOnce, producerCached)
	go pbProducer.WriteToNsq(bundle, nsqpb.PRTopic)

}

func HandleBBEvent(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get("X-Event-Key") {
	case "repo:push":  RepoPush(w, r)
	case "pullrequest:created",
	     "pullrequest:updated": PullRequest(w, r)
	default:
		ocelog.Log().Errorf("No support for Bitbucket event %s", r.Header.Get("X-Event-Key"))
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
}

// for testing
// irl... use vault
func getCredConfig() models.AdminConfig {
	return models.AdminConfig{
		ConfigId:     "jessishank",
		ClientId:     "QEBYwP5cKAC3ykhau4",
		ClientSecret: "gKY2S3NGnFzJKBtUTGjQKc4UNvQqa2Vb",
		TokenURL:     "https://bitbucket.org/site/oauth2/access_token",
		AcctName:     "jessishank",
	}
}

func GetBuildConfig(repoFullName string, checkoutCommit string) (conf *pb.BuildConfig, err error) {
	cfg := getCredConfig()
	bb := handler.Bitbucket{}
	bb.SetMeUp(&cfg)
	fileBitz, err := bb.GetFile("ocelot.yml", repoFullName, checkoutCommit)
	if err != nil {
		return
	}
	conf = &pb.BuildConfig{}
	if err != nil {
		return
	}
	if err = deserializer.YAMLToProto(fileBitz, conf); err != nil {
		return
	}
	return
}


func main() {
	ocelog.InitializeOcelog(ocelog.GetFlags())
	ocelog.Log().Debug()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
		ocelog.Log().Warning("running on default port :8088")
	}
	// initialize vault on startup, we want to know right away if we don't have the creds we need.
	_ = ocevault.GetInitVault(vaultOnce, vaultCached)
	// same goes for nsq producer
	_ = nsqpb.GetInitProducer(producerOnce, producerCached)

	muxi := mux.NewRouter()
	muxi.HandleFunc("/test", HandleBBEvent).Methods("POST")
	// mux.HandleFunc("/", ViewWebhooks).Methods("GET")

	n := negroni.New(negroni.NewRecovery(), negroni.NewStatic(http.Dir("public")))
	n.Use(negronilogrus.NewCustomMiddleware(ocelog.GetLogLevel(), &log.JSONFormatter{}, "web"))
	n.UseHandler(muxi)
	n.Run(":" + port)
}
