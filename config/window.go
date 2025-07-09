package config

// Window represents a tmux's window configuration.
type Window struct {
	Name   string `toml:"name" yaml:"name"`
	Panes  []Pane `toml:"panes" yaml:"panes"`
	Layout string `toml:"layout" yaml:"layout"`
	Sync   *bool  `toml:"sync" yaml:"sync"`
}
