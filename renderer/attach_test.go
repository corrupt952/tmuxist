package renderer

import (
	"testing"

	"tmuxist/config"
	test_helper "tmuxist/helper/test"
)

func TestAttachRenderer_Render(t *testing.T) {
	r := AttachRenderer{&config.Config{}}

	actual := r.Render()
	expected := "tmux attach-session "
	test_helper.AssertEquals(t, actual, expected)
}

func TestAttachRenderer_Render_WithName(t *testing.T) {
	r := AttachRenderer{&config.Config{Name: "tmuxist"}}

	actual := r.Render()
	expected := "tmux attach-session -t tmuxist"
	test_helper.AssertEquals(t, actual, expected)
}
