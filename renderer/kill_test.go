package renderer

import (
	"testing"

	"tmuxist/config"
	test_helper "tmuxist/helper/test"
)

func TestKillRenderer_Render(t *testing.T) {
	r := KillRenderer{&config.Config{}}

	actual := r.Render()
	expected := "tmux kill-session "
	test_helper.AssertEquals(t, actual, expected)
}

func TestKillRenderer_Render_WithName(t *testing.T) {
	r := KillRenderer{&config.Config{Name: "tmuxist"}}

	actual := r.Render()
	expected := "tmux kill-session -t tmuxist"
	test_helper.AssertEquals(t, actual, expected)
}
