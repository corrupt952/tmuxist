package command

import (
	"flag"
	"os"
	"path/filepath"
	"testing"
)

func TestConfigCommand_SetConfigFlags(t *testing.T) {
	cmd := &ConfigCommand{}
	fs := flag.NewFlagSet("test", flag.ContinueOnError)

	cmd.SetConfigFlags(fs)

	// Check if both flags are defined
	shortFlag := fs.Lookup("f")
	if shortFlag == nil {
		t.Error("expected -f flag to be defined")
	}

	longFlag := fs.Lookup("file")
	if longFlag == nil {
		t.Error("expected --file flag to be defined")
	}

	// Test parsing
	err := fs.Parse([]string{"-f", "/path/to/config.yaml"})
	if err != nil {
		t.Fatalf("failed to parse flags: %v", err)
	}

	if cmd.configFile != "/path/to/config.yaml" {
		t.Errorf("expected configFile to be '/path/to/config.yaml', got '%s'", cmd.configFile)
	}
}

func TestConfigCommand_LoadConfig(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "tmuxist-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a test config file
	configPath := filepath.Join(tmpDir, "test-config.yaml")
	configContent := `name: test-session
root: /tmp
windows:
  - panes:
      - command: "echo hello"`

	err = os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("failed to create test config: %v", err)
	}

	tests := []struct {
		name       string
		configFile string
		setup      func()
		wantErr    bool
		cleanup    func()
	}{
		{
			name:       "load specified config file",
			configFile: configPath,
			wantErr:    false,
		},
		{
			name:       "fail on non-existent specified file",
			configFile: "/non/existent/file.yaml",
			wantErr:    true,
		},
		{
			name:    "fail when no default config exists",
			wantErr: true,
			setup: func() {
				// Change to a directory without config files
				os.Chdir(tmpDir)
			},
			cleanup: func() {
				// Return to original directory
				os.Chdir("..")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			if tt.cleanup != nil {
				defer tt.cleanup()
			}

			cmd := &ConfigCommand{configFile: tt.configFile}
			config, err := cmd.LoadConfig()

			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if config == nil {
					t.Error("expected config but got nil")
				} else if config.Name != "test-session" {
					t.Errorf("expected config name to be 'test-session', got '%s'", config.Name)
				}
			}
		})
	}
}

func TestConfigCommand_resolveConfigPath(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "tmuxist-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a test config file
	configPath := filepath.Join(tmpDir, "test.yaml")
	_, err = os.Create(configPath)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	tests := []struct {
		name       string
		configFile string
		wantErr    bool
	}{
		{
			name:       "resolve specified existing file",
			configFile: configPath,
			wantErr:    false,
		},
		{
			name:       "fail on non-existent specified file",
			configFile: "/non/existent/path.yaml",
			wantErr:    true,
		},
		{
			name:       "fail when no config file specified and no default",
			configFile: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ConfigCommand{configFile: tt.configFile}
			path, err := cmd.resolveConfigPath()

			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if path == "" {
					t.Error("expected path but got empty string")
				}
			}
		})
	}
}
