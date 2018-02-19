package config

type Config struct {
	// Server Config
	Port int `json:"port"`

	// Slack Config

	//SlackToken: Get this from https://api.slack.com/apps/<myapp>/install-on-team?
	SlackToken string `json:"slackToken"`
}
