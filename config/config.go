package config

// Config represents a tmuxist and tmux's session configuration.
type Config struct {
	Name    string   `toml:"name" yaml:"name"`
	Root    string   `toml:"root" yaml:"root"`
	Attach  *bool    `toml:"attach" yaml:"attach"`
	Windows []Window `toml:"windows" yaml:"windows"`
}
