package main

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Config struct {
	IsProductionMode bool
	APIAddress       string
	NATSURL          string
	ClusterID        string
	ClientID         string
	DatabaseURL      string
	CacheFile        string
	DataDir          string
}

func CreateConfig() (*Config, error) {
	c := Config{}

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	err := config.LoadExists("config.default.yml", "config.yml")
	if err != nil {
		return &c, err
	}

	config.LoadOSEnv([]string{}, false)

	err = config.BindStruct("", &c)

	return &c, err
}
