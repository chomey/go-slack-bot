package main

import (
	"fmt"
	"github.com/op/go-logging"
	"io/ioutil"
	"encoding/json"
	"github.com/chomey/go-slack-bot/config"
)

var log = logging.MustGetLogger("go-slack-bot")

func main() {
	log.Info("=====")

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		check(err)
	}

	var config config.Config
	json.Unmarshal(data, &config)

	fmt.Println(fmt.Sprintf("Hello %s.", config.Name))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
