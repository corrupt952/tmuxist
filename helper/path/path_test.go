package path

import (
	"os/user"
	"strings"
	"testing"
)

func TestFullpath(t *testing.T) {
	// Get current user for testing
	usr, err := user.Current()
	if err != nil {
		t.Fatalf("failed to get current user: %v", err)
	}

	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "expand tilde at beginning",
			input:    "~/Documents/project",
			expected: usr.HomeDir + "/Documents/project",
			wantErr:  false,
		},
		{
			name:     "no tilde to expand",
			input:    "/usr/local/bin",
			expected: "/usr/local/bin",
			wantErr:  false,
		},
		{
			name:     "tilde in middle of path",
			input:    "/path/~/file",
			expected: "/path/" + usr.HomeDir + "/file",
			wantErr:  false,
		},
		{
			name:     "empty path",
			input:    "",
			expected: "",
			wantErr:  false,
		},
		{
			name:     "just tilde",
			input:    "~",
			expected: usr.HomeDir,
			wantErr:  false,
		},
		{
			name:     "multiple tildes (only first replaced)",
			input:    "~/path/~/file",
			expected: usr.HomeDir + "/path/~/file",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Fullpath(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("expected '%s', got '%s'", tt.expected, result)
				}
			}
		})
	}
}

func TestFullpath_ReplaceOnlyFirstTilde(t *testing.T) {
	// Test specifically that strings.Replace with count=1 only replaces first occurrence
	usr, err := user.Current()
	if err != nil {
		t.Fatalf("failed to get current user: %v", err)
	}

	input := "~~~"
	expected := usr.HomeDir + "~~"

	result, err := Fullpath(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}

	// Verify behavior matches strings.Replace with count=1
	directReplace := strings.Replace(input, "~", usr.HomeDir, 1)
	if result != directReplace {
		t.Errorf("Fullpath result doesn't match strings.Replace behavior")
	}
}
