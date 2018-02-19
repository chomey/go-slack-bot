package main

import (
	"fmt"
	"github.com/op/go-logging"
	"net/http"
	"github.com/chomey/go-slack-bot/config"
	"io/ioutil"
	"encoding/json"
	"github.com/chomey/go-slack-bot/service"
)

type RequestHandler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

var log = logging.MustGetLogger("go_slack_bot")
var Config config.Config
var Service *service.Service
var requestHandlers = make(map[string]RequestHandler)

func main() {
	loadConfig()

	Service = service.New()

	fmt.Printf("Loading config: %#v\n", Config)

	//Register new handlers here
	requestHandlers["/"] = Service
	requestHandlers["/slack"] = Service.Slack

	for path, handler := range requestHandlers {
		http.HandleFunc(path, handler.HandleRequest)
	}

	log.Infof("Now listening on http://localhost:%d\n", Config.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", Config.Port), nil)
	check(err)
}

func loadConfig() {
	data, err := ioutil.ReadFile("config.json")
	check(err)

	err = json.Unmarshal(data, &Config)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
