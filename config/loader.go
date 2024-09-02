package config

import (
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml/v2"

	"tmuxist/logger"
)

// LoadFile returns *config.Config by path.
func LoadFile(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := Config{}
	toml.Unmarshal([]byte(f), &c)
	return &c, nil
}

// ConfigurationPath returns configuration path by profile.
func ConfigurationPath() (string, error) {
	p, err := os.Getwd()
	if err != nil {
		logger.Err(err.Error())
		return "", err
	}

	p = p + "/tmuxist.toml"

	return p, nil
}
