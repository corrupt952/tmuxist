package config

// Config represents a tmuxist and tmux's session configuration.
type Config struct {
	Name    string
	Root    string
	Attach  *bool
	Windows []Window
}
