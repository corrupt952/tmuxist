package shell

import (
	"os"
	"testing"

	test_helper "github.com/corrupt952/tmuxist/helper/test"
)

func TestCommandSubstitution(t *testing.T) {
	os.Setenv("SHELL", "/bin/zsh")

	actual := CommandSubstitution("tmux split-window")
	expected := "$(tmux split-window)"
	test_helper.AssertEquals(t, actual, expected)
}

func TestCommandSubstitution_OnSh(t *testing.T) {
	os.Setenv("SHELL", "/bin/sh")

	actual := CommandSubstitution("tmux split-window")
	expected := "`tmux split-window`"
	test_helper.AssertEquals(t, actual, expected)
}
