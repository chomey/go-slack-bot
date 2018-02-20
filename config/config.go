package config

type Config struct {
	// Server Config
	Port int `json:"port"`

	//
	SlackDebugLogging bool `json:"slackDebugLogging"`
}
