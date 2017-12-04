package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/namsral/flag"
	"github.com/shankj3/ocelot/admin/handler"
	"github.com/shankj3/ocelot/admin/models"
	"github.com/shankj3/ocelot/util/deserialize"
	"github.com/shankj3/ocelot/util/ocelog"
	"github.com/shankj3/ocelot/util/ocenet"
	"io/ioutil"
	"net/http"
	"github.com/shankj3/ocelot/util"
	"gopkg.in/go-playground/validator.v9"
	"context"
)

//TODO: look into hookhandler logic and separate into new ocelot.yaml + new commit
//TODO: rewrite admin code to use grpc
//TODO: floe integration??? just putting this note here so we remember
//TODO: change this to use my fork of logrus so we can pretty print logs


//this is our grpc server struct
type guideOcelotServer struct {
	//TODO: probably want other properties here
}

//TODO: what is htis context good for?
func (g *guideOcelotServer) GetCreds(ctx context.Context) (*models.CredWrapper, error) {
	return nil, nil
}

//application context, contains stuff that'll get used across admin code
type AdminCtx struct {
	deserializer	*deserialize.Deserializer
	adminValidator	*AdminValidator
	remoteConfig	*util.RemoteConfig
}

func main() {
	//load properties
	var port string
	var consulHost string
	var consulPort int
	var logLevel string

	flag.StringVar(&port, "port", "8080", "admin server port")
	flag.StringVar(&consulHost, "consul-host", "localhost", "consul host")
	flag.IntVar(&consulPort, "consul-port", 8500, "consul port")
	flag.StringVar(&logLevel, "log-level", "debug", "ocelot admin log level")
	flag.Parse()

	ocelog.InitializeOcelog(logLevel)

	//TODO: this is my local vault root token, too lazy to set env variable
	configInstance, err := util.GetInstance(consulHost, consulPort, "0fd3c0f4-52ec-7d3b-d29b-4b4df1326ded")

	if err != nil {
		ocelog.Log().Fatal("could not talk to consul or vault, bailing")
	}

	//populate admin context
	adminContext := &AdminCtx{
		deserializer: deserialize.New(),
		adminValidator: GetValidator(),
		remoteConfig: configInstance,
	}

	adminContext.remoteConfig.Consul.RegisterService("localhost", 8080, "ocelot-admin")

	//check for config on load
	ReadConfig(adminContext)

	//start http server
	mux := mux.NewRouter()
	//TODO: seems like maybe this should be command line tool instead - wait for Abby
	//list all configs
	//list all repos + 'tracked' repos vs. 'untracked' repos
	//add new repo
	//configure whether or not you want admin to discover new ocelot.yaml files for you

	mux.Handle("/", &ocenet.AppContextHandler{adminContext, ConfigHandler}).Methods("POST")
	mux.Handle("/", &ocenet.AppContextHandler{adminContext,ListConfigHandler}).Methods("GET")
	ocelog.Log().Fatal(http.ListenAndServe(":"+port, mux))
}

func ListConfigHandler(ctx interface{}, w http.ResponseWriter, r *http.Request) {
	adminCtx := ctx.(*AdminCtx)

	w.Header().Set("Content-Type", "application/json")
	creds, _ := adminCtx.remoteConfig.GetCredAt(util.ConfigPath, true)
	json.NewEncoder(w).Encode(creds)
}

//ConfigHandler handles config from REST api
func ConfigHandler(ctx interface{}, w http.ResponseWriter, r *http.Request) {
	adminCtx := ctx.(*AdminCtx)

	var adminConfig models.AdminConfig
	_ = json.NewDecoder(r.Body).Decode(&adminConfig)

	errorMsg, err := adminCtx.adminValidator.ValidateConfig(&adminConfig)

	if err != nil {
		ocenet.JSONApiError(w, http.StatusBadRequest, errorMsg, err)
		return
	}

	errorMsg, err = SetupCredentials(adminCtx, &adminConfig)
	if err != nil {
		ocenet.JSONApiError(w, http.StatusUnprocessableEntity, errorMsg, err)
		return
	}

}

//reads config file in current directory if it exists, exits if file is unparseable or doesn't exist
func ReadConfig(ctx *AdminCtx) {
	config := &models.ConfigYaml{}
	configFile, err := ioutil.ReadFile(models.ConfigFileName)
	if err != nil {
		ocelog.IncludeErrField(err).Error()
		return
	}
	err = ctx.deserializer.YAMLToStruct(configFile, config)
	if err != nil {
		ocelog.IncludeErrField(err).Error()
		return
	}
	for _, configVal := range config.Credentials {
		errMsg, err := ctx.adminValidator.ValidateConfig(&configVal)
		if err != nil {
			ocelog.IncludeErrField(err).Error(errMsg)
			continue
		}

		_, err = SetupCredentials(ctx, &configVal)
		if err != nil {
			ocelog.IncludeErrField(err).Error()
		}
	}
}

//when new configurations are added to the config channel, create bitbucket client and webhooks
func SetupCredentials(ctx *AdminCtx, config *models.AdminConfig) (string, error) {
	//hehe right now we only have bitbucket
	switch config.Type {
	case "bitbucket":
		bbHandler := handler.Bitbucket{}
		bitbucketClient := &ocenet.OAuthClient{}
		bitbucketClient.Setup(config)

		err := bbHandler.SetMeUp(config, bitbucketClient)

		if err != nil {
			ocelog.Log().Error("could not setup bitbucket client")
			return "Could not setup bitbucket client for " + config.Type + "/" + config.AcctName, err
		}

		err = bbHandler.Walk()
		if err != nil {
			return "Could not traverse repositories and create necessary webhooks for " + config.Type + "/" + config.AcctName, err
		}
	}
	configPath := util.ConfigPath + "/" + config.Type + "/" + config.AcctName
	err := ctx.remoteConfig.AddCreds(configPath, config)
	return "", err
}


////everything below this is for validating/////

func GetValidator() *AdminValidator {
	adminValidator := &AdminValidator {
		Validate: validator.New(),
	}
	adminValidator.Validate.RegisterValidation("validtype", typeValidation)
	return adminValidator
}


//validator for all admin related stuff
type AdminValidator struct {
	Validate	*validator.Validate
}

//validates config and returns json formatted error
func(adminValidator AdminValidator) ValidateConfig(adminConfig *models.AdminConfig) (string, error) {
	err := adminValidator.Validate.Struct(adminConfig)
	if err != nil {
		var errorMsg string
		for _, nestedErr := range err.(validator.ValidationErrors) {
			errorMsg = nestedErr.Field() + " is " + nestedErr.Tag()
			if nestedErr.Tag() == "validtype" {
				errorMsg = "type must be one of the following: bitbucket"
			}

			ocelog.Log().Warn(errorMsg)
		}
		return errorMsg, err
	}
	return "", nil
}

func typeValidation(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case "bitbucket":
		return true
	}
	return false
}