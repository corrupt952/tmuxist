package renderer

import (
	"fmt"

	"tmuxist/config"
)

// AttachRenderer represents startup shell script.
type AttachRenderer struct {
	Config *config.Config
}

// Render returns startup shell script.
func (r *AttachRenderer) Render() string {
	s := ""
	c := r.Config

	name := ""
	if c.Name != "" {
		name = fmt.Sprintf("-t %s", c.Name)
	}
	s += fmt.Sprintf(fmt.Sprintf("tmux attach-session %s", name))

	return s
}
