package command

import (
	"context"
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/subcommands"
)

func TestInitCommand_Name(t *testing.T) {
	cmd := &InitCommand{}
	if cmd.Name() != "init" {
		t.Errorf("expected 'init', got '%s'", cmd.Name())
	}
}

func TestInitCommand_Synopsis(t *testing.T) {
	cmd := &InitCommand{}
	if cmd.Synopsis() != "initialize tmuxist configuration" {
		t.Errorf("expected 'initialize tmuxist configuration', got '%s'", cmd.Synopsis())
	}
}

func TestInitCommand_Usage(t *testing.T) {
	cmd := &InitCommand{}
	expected := "init: tmuxist init\n"
	if cmd.Usage() != expected {
		t.Errorf("expected '%s', got '%s'", expected, cmd.Usage())
	}
}

func TestInitCommand_SetFlags(t *testing.T) {
	cmd := &InitCommand{}
	fs := flag.NewFlagSet("test", flag.ContinueOnError)

	cmd.SetFlags(fs)

	// Check if format flag is defined
	formatFlag := fs.Lookup("format")
	if formatFlag == nil {
		t.Error("expected format flag to be defined")
	}

	// Test default value
	if cmd.format != "yaml" {
		t.Errorf("expected default format to be 'yaml', got '%s'", cmd.format)
	}

	// Test parsing custom format
	err := fs.Parse([]string{"-format", "toml"})
	if err != nil {
		t.Fatalf("failed to parse flags: %v", err)
	}

	if cmd.format != "toml" {
		t.Errorf("expected format to be 'toml', got '%s'", cmd.format)
	}
}

func TestInitCommand_Execute(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "tmuxist-init-test")
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

	tests := []struct {
		name           string
		format         string
		expectedFile   string
		existingFile   string
		wantExitStatus subcommands.ExitStatus
	}{
		{
			name:           "create yaml config",
			format:         "yaml",
			expectedFile:   ".tmuxist.yaml",
			wantExitStatus: subcommands.ExitSuccess,
		},
		{
			name:           "create yml config",
			format:         "yml",
			expectedFile:   ".tmuxist.yml",
			wantExitStatus: subcommands.ExitSuccess,
		},
		{
			name:           "create toml config",
			format:         "toml",
			expectedFile:   ".tmuxist.toml",
			wantExitStatus: subcommands.ExitSuccess,
		},
		{
			name:           "fail when yaml config exists",
			format:         "yaml",
			existingFile:   ".tmuxist.yaml",
			wantExitStatus: subcommands.ExitFailure,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up any existing files
			os.Remove(".tmuxist.yaml")
			os.Remove(".tmuxist.yml")
			os.Remove(".tmuxist.toml")

			// Create existing file if needed
			if tt.existingFile != "" {
				err := os.WriteFile(tt.existingFile, []byte("existing"), 0644)
				if err != nil {
					t.Fatalf("failed to create existing file: %v", err)
				}
			}

			cmd := &InitCommand{format: tt.format}
			status := cmd.Execute(context.Background(), nil)

			if status != tt.wantExitStatus {
				t.Errorf("expected exit status %v, got %v", tt.wantExitStatus, status)
			}

			if tt.wantExitStatus == subcommands.ExitSuccess {
				// Check if file was created
				if _, err := os.Stat(tt.expectedFile); os.IsNotExist(err) {
					t.Errorf("expected file '%s' to be created", tt.expectedFile)
				} else {
					// Verify content
					content, err := os.ReadFile(tt.expectedFile)
					if err != nil {
						t.Fatalf("failed to read created file: %v", err)
					}

					// Check that it contains expected content
					contentStr := string(content)
					if !contains(contentStr, "name = ") && !contains(contentStr, "name:") {
						t.Error("config file doesn't contain expected name field")
					}
					if !contains(contentStr, filepath.Base(tmpDir)) {
						t.Error("config file doesn't contain directory name")
					}
				}
			}
		})
	}
}

func TestInitCommand_Execute_InvalidFormat(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "tmuxist-init-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current dir: %v", err)
	}
	defer os.Chdir(originalDir)

	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("failed to change to temp dir: %v", err)
	}

	// Test with invalid format
	cmd := &InitCommand{format: "invalid"}
	status := cmd.Execute(context.Background(), nil)

	if status != subcommands.ExitFailure {
		t.Errorf("expected failure for invalid format, got %v", status)
	}
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
