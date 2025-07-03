package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"

	"tmuxist/logger"
)

// LoadFile returns *config.Config by path.
func LoadFile(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := Config{}
	ext := strings.ToLower(filepath.Ext(path))
	
	switch ext {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(f, &c)
		if err != nil {
			return nil, err
		}
	case ".toml":
		err = toml.Unmarshal(f, &c)
		if err != nil {
			return nil, err
		}
	default:
		err = toml.Unmarshal(f, &c)
		if err != nil {
			return nil, err
		}
	}
	
	return &c, nil
}

// ConfigurationPath returns configuration path by profile.
func ConfigurationPath() (string, error) {
	p, err := os.Getwd()
	if err != nil {
		logger.Err(err.Error())
		return "", err
	}

	// Check for configuration files in order of preference
	configFiles := []string{
		"tmuxist.yaml",
		".tmuxist.yaml",
		"tmuxist.yml",
		".tmuxist.yml",
		"tmuxist.toml",
		".tmuxist.toml",
	}

	for _, configFile := range configFiles {
		configPath := filepath.Join(p, configFile)
		if _, err := os.Stat(configPath); err == nil {
			return configPath, nil
		}
	}

	// Default to tmuxist.toml if no config file exists
	return filepath.Join(p, "tmuxist.toml"), nil
}
