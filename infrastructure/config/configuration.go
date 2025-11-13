package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type AppConfiguration struct {
	Environment string `envconfig:"APP_ENVIRONMENT" json:"Environment"`
	Mode        string `envconfig:"RUN_MODE" json:"Mode"`

	Server struct {
		Port string `envconfig:"GIN_PORT" json:"Port"`
	} `json:"Server"`

	Session struct {
		WebSockets struct {
			WriteBufferSize int `envconfig:"WEBSOCKETS_WRITE_BUFFER_SIZE" json:"WriteBufferSize"`
			ReadBufferSize  int `envconfig:"WEBSOCKETS_READ_BUFFER_SIZE" json:"ReadBufferSize"`
		} `json:"WebSockets"`
	} `json:"Session"`

	WebRTC struct {
		STUNServers []string `envconfig:"WEBRTC_STUN_SERVERS" json:"STUNServers"`
	} `json:"WebRTC"`
}

func ParseConfiguration() *AppConfiguration {
	var config *AppConfiguration = &AppConfiguration{}
	if err := envconfig.Process("", config); err != nil {
		panic(fmt.Sprintf("Failed to load configuration: \n%s\n", err.Error()))
	}
	fmt.Println(config)
	return config
}
