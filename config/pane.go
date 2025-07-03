package config

// Pane represents a tmux's pane configuration.
type Pane struct {
	Command string `toml:"command" yaml:"command"`
	Size    string `toml:"size" yaml:"size"` // Size in percentage (e.g., "30%")
}
