package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("SHELL", "/bin/zsh")
	code := m.Run()
	os.Exit(code)
}

func TestConfig_ToScript(t *testing.T) {
	c := Config{}

	actual := c.ToScript()
	expected := "SESSION_NO=$(tmux new-session -dP )\n\ntmux attach-session -t $SESSION_NO\n"
	AssertEquals(t, actual, expected)
}

func TestConfig_ToScript_WithRoot(t *testing.T) {
	c := Config{Root: "/var/lib/app"}

	actual := c.ToScript()
	expected := "cd /var/lib/app\nSESSION_NO=$(tmux new-session -dP )\n\ntmux attach-session -t $SESSION_NO\n"
	AssertEquals(t, actual, expected)
}

func TestConfig_ToScript_WithName(t *testing.T) {
	c := Config{Name: "tmuxist"}

	actual := c.ToScript()
	expected := "SESSION_NO=$(tmux new-session -dP -s tmuxist)\n\ntmux attach-session -t $SESSION_NO\n"
	AssertEquals(t, actual, expected)
}

func TestConfig_ToScript_IsNotAttach(t *testing.T) {
	attach := false
	c := Config{Attach: &attach}

	actual := c.ToScript()
	expected := "SESSION_NO=$(tmux new-session -dP )\n\necho $SESSION_NO\n"
	AssertEquals(t, actual, expected)
}

func TestWindow_ToScript(t *testing.T) {
	w := Window{}

	actual := w.ToScript(true)
	expected := "WINDOW_NO=$SESSION_NO\n\n"
	AssertEquals(t, actual, expected)
}

func TestWindow_ToScript_IsNotFirst(t *testing.T) {
	w := Window{}

	actual := w.ToScript(false)
	expected := "WINDOW_NO=$(tmux new-window -t $SESSION_NO -a -P)\n\n"
	AssertEquals(t, actual, expected)
}

func TestWindow_ToScript_WithLayout(t *testing.T) {
	w := Window{Layout: "tiled"}

	actual := w.ToScript(true)
	expected := "WINDOW_NO=$SESSION_NO\ntmux select-layout tiled\n\n"
	AssertEquals(t, actual, expected)
}

func TestWindow_ToScript_SynchronizePanes(t *testing.T) {
	sync := true
	w := Window{Sync: &sync}

	actual := w.ToScript(true)
	expected := "WINDOW_NO=$SESSION_NO\ntmux set sync on\n\n"
	AssertEquals(t, actual, expected)
}

func TestPane_ToScript(t *testing.T) {
	p := Pane{}

	actual := p.ToScript(true)
	expected := "PANE_NO=$WINDOW_NO\n"
	AssertEquals(t, actual, expected)
}

func TestPane_ToScript_IsNotFirst(t *testing.T) {
	p := Pane{}

	actual := p.ToScript(false)
	expected := "PANE_NO=$(tmux split-window -t $WINDOW_NO -P)\n"
	AssertEquals(t, actual, expected)
}

func TestPane_ToScript_WithCommand(t *testing.T) {
	command := `
cd ~

echo "hoge"
`
	p := Pane{command}

	actual := p.ToScript(true)
	expected := `PANE_NO=$WINDOW_NO
tmux send-keys -t $PANE_NO 'cd ~' C-m
tmux send-keys -t $PANE_NO 'echo "hoge"' C-m
`
	AssertEquals(t, actual, expected)
}

func TestCommandSubstitution(t *testing.T) {
	os.Setenv("SHELL", "/bin/zsh")

	actual := commandSubstitution("tmux split-window")
	expected := "$(tmux split-window)"
	AssertEquals(t, actual, expected)
}

func TestCommandSubstitution_OnSh(t *testing.T) {
	os.Setenv("SHELL", "/bin/sh")

	actual := commandSubstitution("tmux split-window")
	expected := "`tmux split-window`"
	AssertEquals(t, actual, expected)
}
