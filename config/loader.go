package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"

	path_helper "github.com/corrupt952/tmuxist/helper/path"
	"github.com/corrupt952/tmuxist/logger"
)

const (
	ConfigDirPath = "~/.config/tmuxist"
)

func LoadFile(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := Config{}
	toml.Unmarshal([]byte(f), &c)
	return &c, nil
}

func LoadFileByProfile(profile string) (*Config, error) {
	p, err := ConfigurationPath(profile)
	if err != nil {
		logger.Err(err.Error())
		return nil, err
	}

	c, err := LoadFile(p)
	if err != nil {
		logger.Err(err.Error())
		return nil, err
	}

	return c, nil
}

func ConfigurationPath(profile string) (string, error) {
	p, err := path_helper.AbsolutePath(filepath.Join(ConfigDirPath, profile+".toml"))
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(p); err != nil {
		return "", err
	}

	return p, nil
}
