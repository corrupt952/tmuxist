package renderer

import (
	"os"
	"strings"
	"testing"

	"tmuxist/config"
	test_helper "tmuxist/helper/test"
)

func TestMain(m *testing.M) {
	os.Setenv("SHELL", "/bin/zsh")
	code := m.Run()
	os.Exit(code)
}

func TestStartRenderer_Render(t *testing.T) {
	r := StartRenderer{&config.Config{}}

	actual := r.Render()
	expected := "SESSION_NO=$(tmux new-session -dP )\n\ntmux attach-session -t $SESSION_NO\n"
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_Render_WithRoot(t *testing.T) {
	r := StartRenderer{&config.Config{Root: "/var/lib/app"}}

	actual := r.Render()
	expected := "cd /var/lib/app\nSESSION_NO=$(tmux new-session -dP )\n\ntmux attach-session -t $SESSION_NO\n"
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_Render_WithName(t *testing.T) {
	r := StartRenderer{&config.Config{Name: "tmuxist"}}

	actual := r.Render()
	expected := "SESSION_NO=$(tmux new-session -dP -s tmuxist)\n\ntmux attach-session -t $SESSION_NO\n"
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_Render_WithEnv(t *testing.T) {
	r := StartRenderer{&config.Config{
		Name: "test-session",
		Env: map[string]string{
			"NODE_ENV": "development",
			"PORT":     "3000",
		},
	}}

	actual := r.Render()
	// Check that environment variables are included in new-session command
	// Note: map iteration order is not guaranteed
	containsEnv1 := strings.Contains(actual, "-e NODE_ENV=development") && strings.Contains(actual, "-e PORT=3000")
	containsEnv2 := strings.Contains(actual, "-e PORT=3000") && strings.Contains(actual, "-e NODE_ENV=development")
	
	if !containsEnv1 && !containsEnv2 {
		t.Errorf("Expected output to contain environment variables, but got: %v", actual)
	}
}

func TestStartRenderer_Render_IsNotAttach(t *testing.T) {
	attach := false
	r := StartRenderer{&config.Config{Attach: &attach}}

	actual := r.Render()
	expected := "SESSION_NO=$(tmux new-session -dP )\n\necho $SESSION_NO\n"
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_RenderWindow(t *testing.T) {
	r := StartRenderer{&config.Config{}}
	w := config.Window{}

	actual := r.renderWindow(&w, true)
	expected := "WINDOW_NO=$SESSION_NO\n\n"
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_RenderWindow_IsNotFirst(t *testing.T) {
	r := StartRenderer{&config.Config{}}
	w := config.Window{}

	actual := r.renderWindow(&w, false)
	expected := "WINDOW_NO=$(tmux new-window -t $SESSION_NO -a -P)\n\n"
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_RenderWindow_WithLayout(t *testing.T) {
	r := StartRenderer{&config.Config{}}
	w := config.Window{Layout: "tiled"}

	actual := r.renderWindow(&w, true)
	expected := "WINDOW_NO=$SESSION_NO\ntmux select-layout tiled\n\n"
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_RenderWindow_WithName(t *testing.T) {
	r := StartRenderer{&config.Config{}}
	w := config.Window{Name: "Development"}

	actual := r.renderWindow(&w, true)
	expected := "WINDOW_NO=$SESSION_NO\ntmux rename-window -t $WINDOW_NO Development\n\n"
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_RenderWindow_WithName_IsNotFirst(t *testing.T) {
	r := StartRenderer{&config.Config{}}
	w := config.Window{Name: "Testing"}

	actual := r.renderWindow(&w, false)
	expected := "WINDOW_NO=$(tmux new-window -t $SESSION_NO -a -P -n Testing)\n\n"
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_RenderWindow_SynchronizePanes(t *testing.T) {
	r := StartRenderer{&config.Config{}}
	sync := true
	w := config.Window{Sync: &sync}

	actual := r.renderWindow(&w, true)
	expected := "WINDOW_NO=$SESSION_NO\ntmux set sync on\n\n"
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_RenderPane(t *testing.T) {
	r := StartRenderer{&config.Config{}}
	p := config.Pane{}

	actual := r.renderPane(&p, true)
	expected := "PANE_NO=$WINDOW_NO\n"
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_RenderPane_IsNotFirst(t *testing.T) {
	r := StartRenderer{&config.Config{}}
	p := config.Pane{}

	actual := r.renderPane(&p, false)
	expected := "PANE_NO=$(tmux split-window -t $WINDOW_NO -P)\n"
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_RenderPane_WithCommand(t *testing.T) {
	r := StartRenderer{&config.Config{}}
	command := `
cd ~

echo "hoge"
`
	p := config.Pane{Command: command}

	actual := r.renderPane(&p, true)
	expected := `PANE_NO=$WINDOW_NO
tmux send-keys -t $PANE_NO 'cd ~' C-m
tmux send-keys -t $PANE_NO 'echo "hoge"' C-m
`
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_RenderWindow_WithGridLayout(t *testing.T) {
	r := StartRenderer{&config.Config{}}
	w := config.Window{
		Layout: "2x2",
		Panes: []config.Pane{
			{Command: "vim"},
			{Command: "npm run dev"},
			{Command: "git status"},
			{Command: "htop"},
		},
	}

	actual := r.renderWindow(&w, true)
	expected := `WINDOW_NO=$SESSION_NO
PANE_NO=$WINDOW_NO
tmux send-keys -t $PANE_NO vim C-m
PANE_NO=$(tmux split-window -t $WINDOW_NO -P)
tmux send-keys -t $PANE_NO 'npm run dev' C-m
PANE_NO=$(tmux split-window -t $WINDOW_NO -P)
tmux send-keys -t $PANE_NO 'git status' C-m
PANE_NO=$(tmux split-window -t $WINDOW_NO -P)
tmux send-keys -t $PANE_NO htop C-m
tmux select-layout tiled

`
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_RenderWindow_WithStandardLayout(t *testing.T) {
	r := StartRenderer{&config.Config{}}
	w := config.Window{
		Layout: "main-vertical",
		Panes: []config.Pane{
			{Command: "vim"},
			{Command: "npm test"},
		},
	}

	actual := r.renderWindow(&w, true)
	expected := `WINDOW_NO=$SESSION_NO
PANE_NO=$WINDOW_NO
tmux send-keys -t $PANE_NO vim C-m
PANE_NO=$(tmux split-window -t $WINDOW_NO -P)
tmux send-keys -t $PANE_NO 'npm test' C-m
tmux select-layout main-vertical

`
	test_helper.AssertEquals(t, actual, expected)
}

func TestStartRenderer_RenderPane_WithSize(t *testing.T) {
	r := StartRenderer{&config.Config{}}
	p := config.Pane{
		Command: "htop",
		Size:    "30%",
	}

	actual := r.renderPane(&p, false)
	expected := `PANE_NO=$(tmux split-window -t $WINDOW_NO -P -p 30)
tmux send-keys -t $PANE_NO htop C-m
`
	test_helper.AssertEquals(t, actual, expected)
}


func TestStartRenderer_RenderWindow_WithPaneSizes(t *testing.T) {
	r := StartRenderer{&config.Config{}}
	w := config.Window{
		Layout: "main-vertical",
		Panes: []config.Pane{
			{Command: "vim"},                   // First pane - no size
			{Command: "npm test", Size: "30%"}, // Second pane - 30%
			{Command: "git log", Size: "20%"},  // Third pane - 20%
		},
	}

	actual := r.renderWindow(&w, true)
	expected := `WINDOW_NO=$SESSION_NO
PANE_NO=$WINDOW_NO
tmux send-keys -t $PANE_NO vim C-m
PANE_NO=$(tmux split-window -t $WINDOW_NO -P -p 30)
tmux send-keys -t $PANE_NO 'npm test' C-m
PANE_NO=$(tmux split-window -t $WINDOW_NO -P -p 20)
tmux send-keys -t $PANE_NO 'git log' C-m
tmux select-layout main-vertical

`
	test_helper.AssertEquals(t, actual, expected)
}
