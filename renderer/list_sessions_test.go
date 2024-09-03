package renderer

import (
	"testing"

	"tmuxist/config"
	test_helper "tmuxist/helper/test"
)

func TestListSessionsRenderer_Render(t *testing.T) {
	r := ListSessionsRenderer{}

	actual := r.Render()
	expected := "tmux list-sessions -F '#{session_name}'"
	test_helper.AssertEquals(t, actual, expected)
}
