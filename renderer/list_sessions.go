package renderer

// ListSessionsRenderer represents startup shell script.
type ListSessionsRenderer struct {}

// Render returns
func (r *ListSessionsRenderer) Render() string {
	return "tmux list-sessions -F '#{session_name}'"
}
