package renderer

import (
	"strings"
	"testing"
	"tmuxist/config"
)

func TestStartRenderer_FullExample_WithGridLayout(t *testing.T) {
	c := &config.Config{
		Name: "test-grid",
		Root: "/tmp",
		Windows: []config.Window{
			{
				Layout: "2x2",
				Panes: []config.Pane{
					{Command: "vim"},
					{Command: "npm run dev"},
					{Command: "git status"},
					{Command: "htop"},
				},
			},
			{
				Layout: "main-horizontal",
				Panes: []config.Pane{
					{Command: "tail -f log.txt"},
					{Command: "watch date"},
				},
			},
		},
	}

	r := StartRenderer{Config: c}
	actual := r.Render()

	// Check that the script contains expected elements
	expectedContains := []string{
		"cd /tmp",
		"tmux new-session -dP -s test-grid",
		"tmux send-keys -t $PANE_NO vim C-m",
		"tmux send-keys -t $PANE_NO 'npm run dev' C-m",
		"tmux send-keys -t $PANE_NO 'git status' C-m",
		"tmux send-keys -t $PANE_NO htop C-m",
		"tmux select-layout tiled",
		"tmux new-window",
		"tmux select-layout main-horizontal",
		"tmux attach-session -t $SESSION_NO",
	}

	for _, expected := range expectedContains {
		if !contains(actual, expected) {
			t.Fatalf("Expected output to contain '%s', but it didn't.\nActual output:\n%s", expected, actual)
		}
	}
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func TestStartRenderer_FullExample_WithPaneSizes(t *testing.T) {
	c := &config.Config{
		Name: "test-sizes",
		Windows: []config.Window{
			{
				Layout: "main-horizontal",
				Panes: []config.Pane{
					{Command: "top"},                 // Main pane
					{Command: "df -h", Size: "25%"},  // 25% of remaining space
					{Command: "ps aux", Size: "30%"}, // 30% of remaining space
				},
			},
		},
	}

	r := StartRenderer{Config: c}
	actual := r.Render()

	// Check that the script contains expected elements
	expectedContains := []string{
		"tmux new-session -dP -s test-sizes",
		"tmux send-keys -t $PANE_NO top C-m",
		"tmux split-window -t $WINDOW_NO -P -p 25",
		"tmux send-keys -t $PANE_NO 'df -h' C-m",
		"tmux split-window -t $WINDOW_NO -P -p 30",
		"tmux send-keys -t $PANE_NO 'ps aux' C-m",
		"tmux select-layout main-horizontal",
		"tmux attach-session -t $SESSION_NO",
	}

	for _, expected := range expectedContains {
		if !contains(actual, expected) {
			t.Fatalf("Expected output to contain '%s', but it didn't.\nActual output:\n%s", expected, actual)
		}
	}
}
