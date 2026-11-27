package renderer

import (
	"testing"

	test_helper "github.com/corrupt952/tmuxist/helper/test"
)

func TestListSessionsRenderer_Render(t *testing.T) {
	r := ListSessionsRenderer{}

	actual := r.Render()
	expected := "tmux list-sessions -F '#{session_name}'"
	test_helper.AssertEquals(t, actual, expected)
}
