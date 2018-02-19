package service

import (
	"github.com/chomey/go-slack-bot/slack"
	"github.com/chomey/go-slack-bot/config"
	"net/http"
	"fmt"
)

type Service struct {
	Slack *slack.SlackClient
}

func New() *Service {
	return &Service{&slack.SlackClient{}}
}

func (s *Service) Init(config config.Config) {
	s.Slack.Init(config)
}

func (s *Service) HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
