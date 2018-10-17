package renderer

import (
	"fmt"
	"strings"

	"github.com/kballard/go-shellquote"

	"github.com/corrupt952/tmuxist/config"
	shell_helper "github.com/corrupt952/tmuxist/helper/shell"
)

// StartRenderer represents startup shell script.
type StartRenderer struct {
	Config *config.Config
}

// Render returns startup shell script.
func (r *StartRenderer) Render() string {
	s := ""
	c := r.Config

	if c.Root != "" {
		s += fmt.Sprintf("cd %s\n", c.Root)
	}

	name := ""
	if c.Name != "" {
		name = fmt.Sprintf("-s %s", c.Name)
	}
	s += fmt.Sprintf("SESSION_NO=%s\n\n", shell_helper.CommandSubstitution(fmt.Sprintf("tmux new-session -dP %s", name)))

	for i, w := range c.Windows {
		s += r.renderWindow(&w, i == 0)
	}

	if c.Attach == nil || *c.Attach {
		s += "tmux attach-session -t $SESSION_NO\n"
	} else {
		s += "echo $SESSION_NO\n"
	}

	return s
}

func (r *StartRenderer) renderWindow(w *config.Window, isFirst bool) string {
	s := ""

	if isFirst {
		s += "WINDOW_NO=$SESSION_NO\n"
	} else {
		s += fmt.Sprintf("WINDOW_NO=%s\n", shell_helper.CommandSubstitution("tmux new-window -t $SESSION_NO -a -P"))
	}

	for i, p := range w.Panes {
		s += r.renderPane(&p, i == 0)
	}

	if w.Layout != "" {
		s += fmt.Sprintf("tmux select-layout %s\n", w.Layout)
	}

	if w.Sync != nil && *w.Sync {
		s += "tmux set sync on\n"
	}

	return s + "\n"
}

func (r *StartRenderer) renderPane(p *config.Pane, isFirst bool) string {
	s := ""

	if isFirst {
		s += "PANE_NO=$WINDOW_NO\n"
	} else {
		s += fmt.Sprintf("PANE_NO=%s\n", shell_helper.CommandSubstitution("tmux split-window -t $WINDOW_NO -P"))
	}

	commands := strings.Split(p.Command, "\n")
	for _, c := range commands {
		if c != "" {
			s += fmt.Sprintf("tmux send-keys -t $PANE_NO %s C-m\n", shellquote.Join(c))
		}
	}

	return s
}
