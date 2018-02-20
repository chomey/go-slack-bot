package slack

import (
	"github.com/nlopes/slack"
	"net/http"
	"fmt"
	"github.com/chomey/go-slack-bot/errorUtils"
	"log"
	"os"
	"github.com/chomey/go-slack-bot/config"
	"encoding/json"
)

const Token = "SLACK_TOKEN"
const BotToken = "SLACK_BOT_TOKEN"

type Slack struct {
	Client *slack.Client
}

func New(variables map[string]string) *Slack {
	slack := &Slack{
		Client: slack.New(variables[BotToken]),
	}
	return slack
}

type MemberJoinedChannel struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Channel     string `json:"channel"`
	ChannelType string `json:"channel_type"`
	Team        string `json:"team"`
	Inviter     string `json:"inviter"`
}

func (s *Slack) Start(Config config.Config) {
	slack.SetLogger(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags))

	rtm := s.Client.NewRTM()
	rtm.SetDebug(Config.SlackDebugLogging)
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Println(msg.Type)
		switch msg.Type {
		case "member_joined_channel":
			memberJoinedChannel := &MemberJoinedChannel{}
			json.Unmarshal(msg.Data.([]byte), memberJoinedChannel )

			fmt.Printf("User: %s joined %s\n", memberJoinedChannel.User, memberJoinedChannel.Channel)
		default:
			//no-op
		}
	}
}

func (s *Slack) HandleRequest(w http.ResponseWriter, r *http.Request) {
	params := slack.PostMessageParameters{

	}
	respChannel, respTimestamp, err := s.Client.PostMessage("#jfoos-bot-testingcfdcf", "This is a test", params)
	errorUtils.Check(err)

	fmt.Fprintf(w, "%s %s", respChannel, respTimestamp)
}
