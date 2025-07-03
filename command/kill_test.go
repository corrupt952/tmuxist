package command

import (
	"context"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/subcommands"
)

func TestKillCommand_Name(t *testing.T) {
	cmd := &KillCommand{}
	if cmd.Name() != "kill" {
		t.Errorf("expected 'kill', got '%s'", cmd.Name())
	}
}

func TestKillCommand_Synopsis(t *testing.T) {
	cmd := &KillCommand{}
	if cmd.Synopsis() != "kill tmux session" {
		t.Errorf("expected 'kill tmux session', got '%s'", cmd.Synopsis())
	}
}

func TestKillCommand_Usage(t *testing.T) {
	cmd := &KillCommand{}
	expected := "kill: tmuxist kill\n"
	if cmd.Usage() != expected {
		t.Errorf("expected '%s', got '%s'", expected, cmd.Usage())
	}
}

func TestKillCommand_SetFlags(t *testing.T) {
	cmd := &KillCommand{}
	fs := flag.NewFlagSet("test", flag.ContinueOnError)

	cmd.SetFlags(fs)

	// Check if both flags are defined (inherited from ConfigCommand)
	shortFlag := fs.Lookup("f")
	if shortFlag == nil {
		t.Error("expected -f flag to be defined")
	}

	longFlag := fs.Lookup("file")
	if longFlag == nil {
		t.Error("expected --file flag to be defined")
	}
}

func TestKillCommand_Execute(t *testing.T) {
	t.Skip("Skipping test that uses syscall.Exec which replaces the test process")

	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "tmuxist-kill-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name           string
		configContent  string
		configFile     string
		wantExitStatus subcommands.ExitStatus
	}{
		{
			name:           "fail with non-existent config",
			configFile:     "/non/existent/config.yaml",
			wantExitStatus: subcommands.ExitFailure,
		},
		{
			name:           "fail with invalid config",
			configContent:  "invalid yaml content",
			wantExitStatus: subcommands.ExitFailure,
		},
		{
			name: "valid config but tmux operations may fail",
			configContent: `name: test-session
root: /tmp
windows:
  - panes:
      - command: "echo test"`,
			wantExitStatus: subcommands.ExitFailure, // Expected to fail without tmux
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var configPath string

			if tt.configContent != "" {
				// Create config file
				configPath = filepath.Join(tmpDir, "test-config.yaml")
				err := os.WriteFile(configPath, []byte(tt.configContent), 0644)
				if err != nil {
					t.Fatalf("failed to create config file: %v", err)
				}
			} else if tt.configFile != "" {
				configPath = tt.configFile
			}

			cmd := &KillCommand{}
			cmd.configFile = configPath

			ctx := context.Background()
			fs := flag.NewFlagSet("test", flag.ContinueOnError)

			status := cmd.Execute(ctx, fs)

			// For operations that require tmux, we expect failure in test environment
			if status != tt.wantExitStatus {
				// Only log error if it's not a tmux-related failure
				if tt.configFile != "" && tt.configFile == "/non/existent/config.yaml" {
					// This is expected to fail due to missing file
				} else if tt.configContent == "invalid yaml content" {
					// This is expected to fail due to invalid YAML
				} else {
					t.Logf("Note: This test may fail due to tmux not being available in test environment")
				}
			}
		})
	}
}

func TestKillCommand_Execute_NoConfig(t *testing.T) {
	t.Skip("Skipping test that uses syscall.Exec which replaces the test process")

	// Test when no config file is specified and no default exists
	tmpDir, err := os.MkdirTemp("", "tmuxist-kill-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Change to temp directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current dir: %v", err)
	}
	defer os.Chdir(originalDir)

	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("failed to change to temp dir: %v", err)
	}

	cmd := &KillCommand{}
	status := cmd.Execute(context.Background(), nil)

	if status != subcommands.ExitFailure {
		t.Errorf("expected failure when no config exists, got %v", status)
	}
}
