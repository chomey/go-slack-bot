package main

import (
	"fmt"
	"github.com/op/go-logging"
	"io/ioutil"
	"encoding/json"
	"github.com/chomey/go-slack-bot/config"
	"net/http"
	"strings"
)

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
	data, err := ioutil.ReadFile("config.json")
	check(err)

	var Config config.Config
	json.Unmarshal(data, &Config)

	fmt.Fprintf(w, "Hello astaxie and %s!", Config.Name) // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName)       // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
