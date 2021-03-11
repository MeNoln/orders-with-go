package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	defaultPort int = 5000
)

// Config struct
type Config struct {
	Port int
	DB   string `yaml:"dbString"`
}

var Cfg *Config

// Load cfg
func Load(filename string) (*Config, error) {
	cfg := Config{
		Port: defaultPort,
	}

	buffer, err := ioutil.ReadFile(fmt.Sprintf("./config/%s.yaml", filename))
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(buffer, &cfg); err != nil {
		return nil, err
	}

	Cfg = &cfg
	return &cfg, err
}
