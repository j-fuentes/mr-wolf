package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config contains the configuration of the bot.
type Config struct {
	AllowedUsers []int `yaml:"allowed-users"`
}

// Read reads the configuration from a file.
func Read(path string) (*Config, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
