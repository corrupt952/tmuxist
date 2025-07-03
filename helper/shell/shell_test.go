package shell

import (
	"os"
	"testing"

	test_helper "tmuxist/helper/test"
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

func TestExec(t *testing.T) {
	tests := []struct {
		name    string
		cmd     string
		wantErr bool
	}{
		{
			name:    "successful command",
			cmd:     "echo test",
			wantErr: false,
		},
		{
			name:    "failing command",
			cmd:     "false",
			wantErr: true,
		},
		{
			name:    "non-existent command",
			cmd:     "/non/existent/command",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Exec(tt.cmd)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestExecWithOutput(t *testing.T) {
	tests := []struct {
		name       string
		cmd        string
		wantOutput string
		wantErr    bool
	}{
		{
			name:       "successful command with output",
			cmd:        "echo hello world",
			wantOutput: "hello world",
			wantErr:    false,
		},
		{
			name:       "command with no output",
			cmd:        "true",
			wantOutput: "",
			wantErr:    false,
		},
		{
			name:       "failing command",
			cmd:        "false",
			wantOutput: "",
			wantErr:    true,
		},
		{
			name:       "non-existent command",
			cmd:        "/non/existent/command",
			wantOutput: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := ExecWithOutput(tt.cmd)

			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if output != tt.wantOutput {
					t.Errorf("expected output '%s', got '%s'", tt.wantOutput, output)
				}
			}
		})
	}
}

func TestCurrentShell(t *testing.T) {
	// Test various shell paths
	tests := []struct {
		name     string
		shell    string
		expected string
	}{
		{
			name:     "zsh",
			shell:    "/bin/zsh",
			expected: "/bin/zsh",
		},
		{
			name:     "bash",
			shell:    "/bin/bash",
			expected: "/bin/bash",
		},
		{
			name:     "sh",
			shell:    "/bin/sh",
			expected: "/bin/sh",
		},
		{
			name:     "fish with path",
			shell:    "/usr/local/bin/fish",
			expected: "/usr/local/bin/fish",
		},
		{
			name:     "empty shell",
			shell:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("SHELL", tt.shell)
			result := CurrentShell()
			test_helper.AssertEquals(t, result, tt.expected)
		})
	}
}
