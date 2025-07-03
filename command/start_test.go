package command

import (
	"context"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/subcommands"
)

func TestStartCommand_Name(t *testing.T) {
	cmd := &StartCommand{}
	if cmd.Name() != "start" {
		t.Errorf("expected 'start', got '%s'", cmd.Name())
	}
}

func TestStartCommand_Synopsis(t *testing.T) {
	cmd := &StartCommand{}
	if cmd.Synopsis() != "start tmux session" {
		t.Errorf("expected 'start tmux session', got '%s'", cmd.Synopsis())
	}
}

func TestStartCommand_Usage(t *testing.T) {
	cmd := &StartCommand{}
	expected := "start: tmuxist start\n"
	if cmd.Usage() != expected {
		t.Errorf("expected '%s', got '%s'", expected, cmd.Usage())
	}
}

func TestStartCommand_SetFlags(t *testing.T) {
	cmd := &StartCommand{}
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

func TestStartCommand_Execute(t *testing.T) {
	t.Skip("Skipping test that uses syscall.Exec which replaces the test process")

	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "tmuxist-start-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name           string
		configContent  string
		configFile     string
		wantExitStatus subcommands.ExitStatus
		skipTmuxCheck  bool
	}{
		{
			name:           "fail with non-existent config",
			configFile:     "/non/existent/config.yaml",
			wantExitStatus: subcommands.ExitFailure,
			skipTmuxCheck:  true,
		},
		{
			name:           "fail with invalid config",
			configContent:  "invalid yaml content",
			wantExitStatus: subcommands.ExitFailure,
			skipTmuxCheck:  true,
		},
		{
			name: "valid config but tmux operations may fail",
			configContent: `name: test-session
root: /tmp
windows:
  - panes:
      - command: "echo test"`,
			wantExitStatus: subcommands.ExitFailure, // Expected to fail without tmux
			skipTmuxCheck:  false,
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

			cmd := &StartCommand{}
			cmd.configFile = configPath

			ctx := context.Background()
			fs := flag.NewFlagSet("test", flag.ContinueOnError)

			status := cmd.Execute(ctx, fs)

			// For operations that require tmux, we expect failure in test environment
			if status != tt.wantExitStatus {
				if !tt.skipTmuxCheck {
					t.Logf("Note: This test expects tmux operations to fail in test environment")
				}
				t.Errorf("expected exit status %v, got %v", tt.wantExitStatus, status)
			}
		})
	}
}

func TestStartCommand_Execute_NoConfig(t *testing.T) {
	t.Skip("Skipping test that uses syscall.Exec which replaces the test process")

	// Test when no config file is specified and no default exists
	tmpDir, err := os.MkdirTemp("", "tmuxist-start-test")
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

	cmd := &StartCommand{}
	status := cmd.Execute(context.Background(), nil)

	if status != subcommands.ExitFailure {
		t.Errorf("expected failure when no config exists, got %v", status)
	}
}
