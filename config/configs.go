package config

import (
	"log"
	"sync"
)

type Config struct {
	HttpPort     string `env:"HTTP_PORT" default:"6001"`
	GrpcPort     string `env:"GRPC_PORT" default:"6000"`
	UrlLoginToFM string `env:"URL_LOGIN_TO_FM" default:""`
}

var configInstance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		configInstance = &Config{}
		configInstance.LoadFromEnv()
		configInstance.LoadDefault()
		if err := configInstance.Validate(); err != nil {
			log.Fatalf("Invalid config: %s", err)
		}
	})

	return configInstance
}

func (c *Config) LoadFromEnv() {

}

func (c *Config) LoadDefault() {
	if c.HttpPort == "" {
		c.HttpPort = "6001"
	}
	if c.GrpcPort == "" {
		c.GrpcPort = "6000"
	}
	if c.UrlLoginToFM == "" {
		c.UrlLoginToFM = "http://localhost:4000/auth/token"
	}
}

func (c *Config) Validate() error {
	return nil
}
