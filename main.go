package main

import (
	"fmt"
	"github.com/op/go-logging"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"strings"
	"github.com/chomey/go-slack-bot/config"
)

var Config config.Config

var log = logging.MustGetLogger("go_slack_bot")

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello astaxie and %s!", Config.Name) // send data to client side
}

func main() {
	data, err := ioutil.ReadFile("config.json")
	check(err)

	err = json.Unmarshal(data, &Config)
	check(err)

	fmt.Printf("Loading config: %#v\n", Config)

	err = registerHandlers()
	check(err)
}

func registerHandlers() error {
	http.HandleFunc("/", sayhelloName)

	fmt.Printf("Now listening on http://localhost:%d\n", Config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", Config.Port), nil)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
