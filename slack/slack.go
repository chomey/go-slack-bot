package slack

import (
	"github.com/nlopes/slack"
	"github.com/chomey/go-slack-bot/config"
	"net/http"
	"fmt"
)

type SlackClient struct {
	Client *slack.Client
}

func (s *SlackClient) Init(config config.Config) {
	s.Client = slack.New(config.SlackToken)
}

func (s *SlackClient) HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Slack World!")
}
