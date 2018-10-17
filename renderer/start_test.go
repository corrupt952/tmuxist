package renderer

import (
	"os"
	"testing"

	"github.com/corrupt952/tmuxist/config"
	test_helper "github.com/corrupt952/tmuxist/helper/test"
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
	p := config.Pane{command}

	actual := r.renderPane(&p, true)
	expected := `PANE_NO=$WINDOW_NO
tmux send-keys -t $PANE_NO 'cd ~' C-m
tmux send-keys -t $PANE_NO 'echo "hoge"' C-m
`
	test_helper.AssertEquals(t, actual, expected)
}
