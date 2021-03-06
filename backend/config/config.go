package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Server struct {
	Port int  `yaml:"port"`
	CORS CORS `yaml:"cors"`
}

// Config .
type Config struct {
	Server Server `yaml:"server"`
}

type CORS struct {
	Hosts []string `yaml:"hosts"`
}

func NewConfig(path string) (*Config, error) {
	cfg := new(Config)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
