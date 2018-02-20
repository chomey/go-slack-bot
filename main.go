package main

import (
	"fmt"
	"github.com/op/go-logging"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/chomey/go-slack-bot/slack"
	"github.com/chomey/go-slack-bot/config"
	"github.com/gorilla/mux"
	"os"
	"errors"
	"github.com/chomey/go-slack-bot/errorUtils"
)

type RequestHandler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

var log = logging.MustGetLogger("go_slack_bot")
var Slack *slack.Slack
var requestHandlers = make(map[string]RequestHandler)
var variables = make(map[string]string)

func main() {
	Config := loadConfig()
	fmt.Printf("Loading Config: %#v\n", Config)

	loadEnvironmentVariables()
	fmt.Printf("%#v", variables)

	Slack = slack.New(variables)
	Slack.Start(Config)

	m := setupMuxRouter()
	http.Handle("/", m)
	log.Infof("Now listening on http://localhost:%d\n", Config.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", Config.Port), nil)
	errorUtils.Check(err)
}

func setupMuxRouter() *mux.Router {
	m := mux.NewRouter()
	//Register new handlers here
	requestHandlers["/slack"] = Slack
	requestHandlers["/"] = Slack
	for path, handler := range requestHandlers {
		m.PathPrefix(path).HandlerFunc(handler.HandleRequest)
	}
	return m
}

func loadConfig() config.Config {
	data, err := ioutil.ReadFile("config.json")
	errorUtils.Check(err)

	Config := new(config.Config)
	err = json.Unmarshal(data, &Config)
	errorUtils.Check(err)
	return *Config
}

func loadEnvironmentVariables() {
	checkIfSet(slack.Token)
	checkIfSet(slack.BotToken)
}

func checkIfSet(environmentVariable string) {
	ev := os.Getenv(environmentVariable)
	if ev == "" {
		errorUtils.Check(errors.New(fmt.Sprintf("Environment variable '%s' is requried", environmentVariable)))
	}
	variables[environmentVariable] = ev
}