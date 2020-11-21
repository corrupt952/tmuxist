package renderer

import (
	"fmt"

	"tmuxist/config"
)

// KillRenderer represents startup shell script.
type KillRenderer struct {
	Config *config.Config
}

// Render returns startup shell script.
func (r *KillRenderer) Render() string {
	s := ""
	c := r.Config

	name := ""
	if c.Name != "" {
		name = fmt.Sprintf("-t %s", c.Name)
	}
	s += fmt.Sprintf(fmt.Sprintf("tmux kill-session %s", name))

	return s
}
