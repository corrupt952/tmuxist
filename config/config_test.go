package config

import (
	"testing"

	test_helper "tmuxist/helper/test"
)

func TestSanitizeSessionName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no special characters",
			input:    "my-session",
			expected: "my-session",
		},
		{
			name:     "with period",
			input:    "my.session",
			expected: "my_session",
		},
		{
			name:     "with colon",
			input:    "my:session",
			expected: "my_session",
		},
		{
			name:     "with multiple periods",
			input:    "my.project.dev",
			expected: "my_project_dev",
		},
		{
			name:     "with multiple colons",
			input:    "my:project:dev",
			expected: "my_project_dev",
		},
		{
			name:     "with both period and colon",
			input:    "my.project:dev",
			expected: "my_project_dev",
		},
		{
			name:     "domain name style",
			input:    "example.com",
			expected: "example_com",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only special characters",
			input:    ".:",
			expected: "__",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeSessionName(tt.input)
			test_helper.AssertEquals(t, result, tt.expected)
		})
	}
}
