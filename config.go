package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kballard/go-shellquote"
)

type Config struct {
	Name    string
	Root    string
	Attach  *bool
	Windows []Window
}

func (c *Config) ToScript() string {
	s := ""

	if c.Root != "" {
		s += fmt.Sprintf("cd %s\n", c.Root)
	}

	name := ""
	if c.Name != "" {
		name = fmt.Sprintf("-s %s", c.Name)
	}
	s += fmt.Sprintf("SESSION_NO=%s\n\n", quote(fmt.Sprintf("tmux new-session -dP %s", name)))

	for i, w := range c.Windows {
		s += w.ToScript(i == 0)
	}

	if c.Attach == nil || *c.Attach {
		s += "tmux attach-session -t $SESSION_NO\n"
	} else {
		s += "echo $SESSION_NO\n"
	}

	return s
}

type Window struct {
	Panes []Pane
}

func (w *Window) ToScript(isFirst bool) string {
	s := ""

	if isFirst {
		s += "WINDOW_NO=$SESSION_NO\n"
	} else {
		s += fmt.Sprintf("WINDOW_NO=%s\n", quote("tmux new-window -t $SESSION_NO -a -P"))
	}

	for i, p := range w.Panes {
		s += p.ToScript(i == 0)
	}

	return s + "\n"
}

type Pane struct {
	Command string
}

func (p *Pane) ToScript(isFirst bool) string {
	s := ""

	if isFirst {
		s += "PANE_NO=$WINDOW_NO\n"
	} else {
		s += fmt.Sprintf("PANE_NO=%s\n", quote("tmux split-window -t $WINDOW_NO -P"))
	}

	commands := strings.Split(p.Command, "\n")
	for _, c := range commands {
		if c != "" {
			s += fmt.Sprintf("tmux send-keys -t $PANE_NO %s C-m\n", shellquote.Join(c))
		}
	}

	return s
}

func quote(s string) string {
	shellEnv := os.Getenv("SHELL")
	shell := filepath.Base(shellEnv)

	switch shell {
	case "zsh":
		return fmt.Sprintf("$(%s)", s)
	default:
		return fmt.Sprintf("`%s`", s)
	}
}
