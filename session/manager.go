package session

import (
	"strings"

	shell_helper "tmuxist/helper/shell"
	"tmuxist/renderer"
)

// DefaultSessionLister is the default implementation of SessionLister
type DefaultSessionLister struct{}

// ListSessions returns a list of active tmux sessions
func (d *DefaultSessionLister) ListSessions() ([]string, error) {
	r := renderer.ListSessionsRenderer{}
	output, err := shell_helper.ExecWithOutput(r.Render())
	if err != nil {
		// If tmux server is not running or other errors occur,
		// return empty list instead of error
		return []string{}, nil
	}

	lines := strings.Split(output, "\n")
	var sessions []string
	for _, line := range lines {
		if line != "" {
			sessions = append(sessions, line)
		}
	}

	return sessions, nil
}

// Manager handles tmux session operations
type Manager struct {
	lister SessionLister
}

// NewManager creates a new session manager
func NewManager(lister SessionLister) *Manager {
	if lister == nil {
		lister = &DefaultSessionLister{}
	}
	return &Manager{lister: lister}
}

// HasSession checks if a tmux session with the given name already exists
func (m *Manager) HasSession(name string) (bool, error) {
	sessions, err := m.lister.ListSessions()
	if err != nil {
		return false, err
	}

	for _, session := range sessions {
		if session == name {
			return true, nil
		}
	}

	return false, nil
}

// HasSession is a convenience function using the default lister
func HasSession(name string) (bool, error) {
	manager := NewManager(nil)
	return manager.HasSession(name)
}
