package renderer

import (
	"fmt"
	"strings"

	"github.com/kballard/go-shellquote"

	"tmuxist/config"
	shell_helper "tmuxist/helper/shell"
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
	
	// Add environment variables if specified
	envOpts := ""
	if len(c.Env) > 0 {
		envParts := []string{}
		for k, v := range c.Env {
			envParts = append(envParts, fmt.Sprintf("-e %s=%s", k, shellquote.Join(v)))
		}
		envOpts = " " + strings.Join(envParts, " ")
	}
	
	s += fmt.Sprintf("SESSION_NO=%s\n\n", shell_helper.CommandSubstitution(fmt.Sprintf("tmux new-session -dP %s%s", name, envOpts)))

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
		// Check if it's a grid layout notation (e.g., "2x2")
		if isGridLayout(w.Layout) {
			cols, rows, _ := parseGridLayout(w.Layout)
			s += generateGridCommands(cols, rows)
		} else {
			s += fmt.Sprintf("tmux select-layout %s\n", w.Layout)
		}
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
		splitCmd := "tmux split-window -t $WINDOW_NO -P"
		// Add size option if specified
		if p.Size != "" && strings.HasSuffix(p.Size, "%") {
			// Extract percentage value
			percentage := strings.TrimSuffix(p.Size, "%")
			splitCmd = fmt.Sprintf("tmux split-window -t $WINDOW_NO -P -p %s", percentage)
		}
		s += fmt.Sprintf("PANE_NO=%s\n", shell_helper.CommandSubstitution(splitCmd))
	}

	commands := strings.Split(p.Command, "\n")
	for _, c := range commands {
		if c != "" {
			s += fmt.Sprintf("tmux send-keys -t $PANE_NO %s C-m\n", shellquote.Join(c))
		}
	}

	return s
}
