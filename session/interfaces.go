package session

// SessionLister is an interface for listing tmux sessions
type SessionLister interface {
	ListSessions() ([]string, error)
}
