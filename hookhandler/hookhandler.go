package hookhandler

import (
	"bitbucket.org/level11consulting/go-til/deserialize"
	ocelog "bitbucket.org/level11consulting/go-til/log"
	ocenet "bitbucket.org/level11consulting/go-til/net"
	"bitbucket.org/level11consulting/go-til/nsqpb"
	res "bitbucket.org/level11consulting/leveler_resources"
	"bitbucket.org/level11consulting/ocelot/admin/handler"
	"bitbucket.org/level11consulting/ocelot/admin/models"
	pb "bitbucket.org/level11consulting/ocelot/protos"
	"bitbucket.org/level11consulting/ocelot/util/cred"
	"errors"
	"net/http"
	"strings"
)

type HookHandlerContext struct {
	RemoteConfig *cred.RemoteConfig
	Producer     *nsqpb.PbProduce
	Deserializer *deserialize.Deserializer
}

//TODO: look into all the branches that's listed inside of ocelot.yml and only build if event corresonds
//tODO: branch inside of ocelot.yml
//TODO: what data do we have to store/do we need to store?
// On receive of repo push, marshal the json to an object then build the appropriate pipeline config and put on NSQ queue.
func RepoPush(ctx *HookHandlerContext, w http.ResponseWriter, r *http.Request) {
	repopush := &pb.RepoPush{}
	if err := ctx.Deserializer.JSONToProto(r.Body, repopush); err != nil {
		ocenet.JSONApiError(w, http.StatusBadRequest, "could not parse request body into proto.Message", err)
	}
	fullName := repopush.Repository.FullName
	hash := repopush.Push.Changes[0].New.Target.Hash
	acctName := repopush.Repository.Owner.Username
	buildConf, err := GetBBBuildConfig(ctx, acctName, fullName, hash)
	if err != nil {
		// if the build file just isn't there don't worry about it.
		if err != ocenet.FileNotFound {
			ocelog.IncludeErrField(err).Error("unable to get build conf")
			return
		}
		ocelog.Log().Debugf("no ocelot yml found for repo %s", repopush.Repository.FullName)
		return
	}
	tellWerker(ctx, buildConf, hash)
}

//TODO: look into all the branches that's listed inside of ocelot.yml and only build if event corresonds
//tODO: branch inside of ocelot.yml
// On receive of pull request, marshal the json to an object then build the appropriate pipeline config and put on NSQ queue.
func PullRequest(ctx *HookHandlerContext, w http.ResponseWriter, r *http.Request) {
	pr := &pb.PullRequest{}
	if err := ctx.Deserializer.JSONToProto(r.Body, pr); err != nil {
		ocelog.IncludeErrField(err).Error("could not parse request body into pb.PullRequest")
		return
	}
	fullName := pr.Pullrequest.Source.Repository.FullName
	hash := pr.Pullrequest.Source.Commit.Hash
	acctName := pr.Pullrequest.Source.Repository.Owner.Username

	buildConf, err := GetBBBuildConfig(ctx, acctName, fullName, hash)
	if err != nil {
		// if the build file just isn't there don't worry about it.
		if err != ocenet.FileNotFound {
			ocelog.IncludeErrField(err).Error("unable to get build conf")
			return
		}
		ocelog.Log().Debugf("no ocelot yml found for repo %s", pr.Pullrequest.Source.Repository.FullName)
		return
	}
	tellWerker(ctx, buildConf, hash)
}

func tellWerker(ctx *HookHandlerContext, buildConf *pb.BuildConfig, hash string) {
	// get one-time token use for access to vault
	token, err := ctx.RemoteConfig.Vault.CreateThrowawayToken()
	if err != nil {
		ocelog.IncludeErrField(err).Error("unable to create one-time vault token")
		return
	}

	pipeConfig, err := werk(*buildConf, hash)

	werkerTask := &pb.WerkerTask{
		VaultToken:   token,
		CheckoutHash: hash,
		Pipe:         pipeConfig,
	}

	go ctx.Producer.WriteProto(werkerTask, "docker")
}

//TODO: state = not started = to be stored inside of postgres (db interface is gonna be inside of go-til)
//this just builds the pipeline config, worker will call NewPipeline with the pipeline config and run
func werk(oceConfig pb.BuildConfig, gitCommit string) (*res.PipelineConfig, error) {
	//TODO: example input for job? What should be passed to list of strings?
	// inputs/outputs in a JOB are the keys to pipeline input/outputs in PipelineConfig
	//TODO: how/when do we push artifacts to nexus? (think about this while I'm writing other code)
	// TODO: potentially watch for changes in .m2/PKG_NAME with fsnotify?
	//TODO: we might be able to actually create an image and use input/outputs for the packages part?

	jobMap := make(map[string]*res.JobConfig)

	var kickOffCmd []string
	var kickOffEnvs = make(map[string]string)
	var buildImage string

	if oceConfig.Image != "" {
		buildImage = oceConfig.Image
	} else if len(oceConfig.Packages) > 0 {
		buildImage = "TODO PARSE THIS AND PUSH TO ARTIFACT REPO"
		//TODO: build image and store it somewhere. OH! NEXUS! oh shit we need nexus int. now
	}

	if oceConfig.Before != nil {
		if oceConfig.Before.Script != nil {
			kickOffCmd = append(kickOffCmd, oceConfig.Before.Script...)
		}

		//combine optional before env values if passed
		if oceConfig.Before.Env != nil {
			for envKey, envVal := range oceConfig.Before.Env {
				kickOffEnvs[envKey] = envVal
			}
		}
	}

	if oceConfig.Build != nil {
		if oceConfig.Build.Script != nil {
			kickOffCmd = append(kickOffCmd, oceConfig.Build.Script...)
		}

		//combine optional before env values if passed
		if oceConfig.Build.Env != nil {
			for envKey, envVal := range oceConfig.Build.Env {
				kickOffEnvs[envKey] = envVal
			}
		}
	}

	//TODO: where to store failed builds
	if len(kickOffCmd) == 0 {
		return nil, errors.New("You must have at least one stage populated to trigger a build")
	}

	//create a settings.xml maven file that takes in nexus and/or something else creds

	//TODO: figure out what to do about the rest of the stages
	job := &res.JobConfig{
		Command: strings.Join(kickOffCmd, " && "),
		Env:     kickOffEnvs,
		Image:   buildImage,
	}

	jobMap[gitCommit] = job

	pipeConfig := &res.PipelineConfig{
		Steps:     jobMap,
		GlobalEnv: oceConfig.Env,
	}
	return pipeConfig, nil
}

func HandleBBEvent(ctx interface{}, w http.ResponseWriter, r *http.Request) {
	handlerCtx := ctx.(*HookHandlerContext)

	switch r.Header.Get("X-Event-Key") {
	case "repo:push":
		RepoPush(handlerCtx, w, r)
	case "pullrequest:created",
		"pullrequest:updated":
		PullRequest(handlerCtx, w, r)
	default:
		ocelog.Log().Errorf("No support for Bitbucket event %s", r.Header.Get("X-Event-Key"))
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
}

// for testing
func getCredConfig() *models.Credentials {
	return &models.Credentials{
		ClientId:     "QEBYwP5cKAC3ykhau4",
		ClientSecret: "gKY2S3NGnFzJKBtUTGjQKc4UNvQqa2Vb",
		TokenURL:     "https://bitbucket.org/site/oauth2/access_token",
		AcctName:     "jessishank",
	}
}

func GetBBBuildConfig(ctx *HookHandlerContext, acctName string, repoFullName string, checkoutCommit string) (conf *pb.BuildConfig, err error) {
	//cfg := getCredConfig()
	bbCreds, err := ctx.RemoteConfig.GetCredAt(cred.BuildVCSCredPath("bitbucket", acctName), false)
	cfg := bbCreds["bitbucket/"+acctName]
	bb := handler.Bitbucket{}
	bbClient := &ocenet.OAuthClient{}
	bbClient.Setup(cfg)

	bb.SetMeUp(cfg, bbClient)
	fileBitz, err := bb.GetFile("ocelot.yml", repoFullName, checkoutCommit)
	if err != nil {
		return
	}
	conf = &pb.BuildConfig{}
	if err != nil {
		return
	}
	if err = ctx.Deserializer.YAMLToStruct(fileBitz, conf); err != nil {
		return
	}
	return
}
