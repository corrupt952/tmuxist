package config

import "strings"

// Config represents a tmuxist and tmux's session configuration.
type Config struct {
	Name    string            `toml:"name" yaml:"name"`
	Root    string            `toml:"root" yaml:"root"`
	Attach  *bool             `toml:"attach" yaml:"attach"`
	Env     map[string]string `toml:"env" yaml:"env"`
	Windows []Window          `toml:"windows" yaml:"windows"`
}

// SanitizeSessionName sanitizes a tmux session name by replacing
// characters that tmux doesn't allow (colon and period) with underscores.
// This matches tmux's internal behavior in session_check_name().
func SanitizeSessionName(name string) string {
	name = strings.ReplaceAll(name, ":", "_")
	name = strings.ReplaceAll(name, ".", "_")
	return name
}
