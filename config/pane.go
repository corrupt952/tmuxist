package config

// Pane represents a tmux's pane configuration.
type Pane struct {
	Command string `toml:"command" yaml:"command"`
}
