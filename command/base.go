package command

import (
	"flag"
	"os"

	"tmuxist/config"
)

// ConfigCommand is a base struct for commands that need configuration file
type ConfigCommand struct {
	configFile string
}

// SetConfigFlags adds configuration file flags to the flag set
func (c *ConfigCommand) SetConfigFlags(f *flag.FlagSet) {
	f.StringVar(&c.configFile, "f", "", "Path to the configuration file")
	f.StringVar(&c.configFile, "file", "", "Path to the configuration file")
}

// LoadConfig loads the configuration from file
func (c *ConfigCommand) LoadConfig() (*config.Config, error) {
	path, err := c.resolveConfigPath()
	if err != nil {
		return nil, err
	}
	return config.LoadFile(path)
}

// resolveConfigPath determines the configuration file path
func (c *ConfigCommand) resolveConfigPath() (string, error) {
	if c.configFile != "" {
		// Use the specified file
		if _, err := os.Stat(c.configFile); err != nil {
			return "", err
		}
		return c.configFile, nil
	}

	// Use the default configuration path
	path, err := config.ConfigurationPath()
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(path); err != nil {
		return "", err
	}

	return path, nil
}
