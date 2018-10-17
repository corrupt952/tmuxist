package config

// Window represents a tmux's window configuration.
type Window struct {
	Panes  []Pane
	Layout string
	Sync   *bool
}
