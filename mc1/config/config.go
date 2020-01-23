package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config holds data of configuration file
type Config struct {
	Server *Server `yaml:"server"`
	JWT    *JWT    `yaml:"jwt"`
}

// Load returns a *Config with all the configuration data
// needed to start the server
func Load(environment string) (*Config, error) {
	if environment == "" {
		return nil, fmt.Errorf("environment not specified, got %s", environment)
	}

	fileName := fmt.Sprintf("./config/%s.yml", environment)

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file: %s, %s", fileName, err.Error())
	}

	cfg := &Config{}
	if err = yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to unmarshall configuration: %s", err.Error())
	}

	return cfg, nil
}
