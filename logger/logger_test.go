package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestSetup(t *testing.T) {
	var buf bytes.Buffer
	Setup(&buf)

	// Test that logger is initialized
	if l == nil {
		t.Error("expected logger to be initialized after Setup")
	}

	// Test that Setup can be called multiple times without panic
	Setup(&buf)
}

func TestLogger_logging(t *testing.T) {
	var buf bytes.Buffer
	Setup(&buf)

	tests := []struct {
		name     string
		logFunc  func(string)
		level    string
		message  string
		expected string
	}{
		{
			name:     "debug message",
			logFunc:  Debug,
			level:    LogLevelDebug,
			message:  "debug test message",
			expected: "", // Debug is below min level (WARN)
		},
		{
			name:     "warn message",
			logFunc:  Warn,
			level:    LogLevelWarn,
			message:  "warning test message",
			expected: "[WARN] warning test message",
		},
		{
			name:     "error message",
			logFunc:  Err,
			level:    LogLevelErr,
			message:  "error test message",
			expected: "[ERROR] error test message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()

			tt.logFunc(tt.message)

			output := buf.String()
			if tt.expected == "" {
				if output != "" {
					t.Errorf("expected no output for %s level, but got: %s", tt.level, output)
				}
			} else {
				if !strings.Contains(output, tt.expected) {
					t.Errorf("expected output to contain '%s', but got: %s", tt.expected, output)
				}
			}
		})
	}
}

func TestLogger_levelFilter(t *testing.T) {
	var buf bytes.Buffer
	logger := &Logger{}

	// Test with different min levels
	tests := []struct {
		name      string
		minLevel  string
		testMsg   string
		testLvl   string
		shouldLog bool
	}{
		{
			name:      "debug with warn min level",
			minLevel:  LogLevelWarn,
			testMsg:   "test",
			testLvl:   LogLevelDebug,
			shouldLog: false,
		},
		{
			name:      "warn with warn min level",
			minLevel:  LogLevelWarn,
			testMsg:   "test",
			testLvl:   LogLevelWarn,
			shouldLog: true,
		},
		{
			name:      "error with debug min level",
			minLevel:  LogLevelDebug,
			testMsg:   "test",
			testLvl:   LogLevelErr,
			shouldLog: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			filter := logger.levelFilter(&buf, tt.minLevel)

			// Write a log message through the filter
			filter.Write([]byte("[" + tt.testLvl + "] " + tt.testMsg))

			output := buf.String()
			if tt.shouldLog && output == "" {
				t.Error("expected log output but got none")
			} else if !tt.shouldLog && output != "" {
				t.Errorf("expected no output but got: %s", output)
			}
		})
	}
}

func TestLoggerConstants(t *testing.T) {
	// Test that constants have expected values
	if LogLevelDebug != "DEBUG" {
		t.Errorf("expected LogLevelDebug to be 'DEBUG', got '%s'", LogLevelDebug)
	}
	if LogLevelWarn != "WARN" {
		t.Errorf("expected LogLevelWarn to be 'WARN', got '%s'", LogLevelWarn)
	}
	if LogLevelErr != "ERROR" {
		t.Errorf("expected LogLevelErr to be 'ERROR', got '%s'", LogLevelErr)
	}
}

func TestLoggerWithCustomMinLevel(t *testing.T) {
	// Test changing log level by creating custom filter
	var buf bytes.Buffer
	logger := &Logger{}
	l = logger // Set global logger

	// Set up with DEBUG level to see all messages
	filter := logger.levelFilter(&buf, LogLevelDebug)
	// Note: In real usage, we'd need to set log.SetOutput(filter)
	// But for testing, we'll directly use the filter

	// Test that all levels are included in the filter
	levels := filter.Levels
	if len(levels) != 3 {
		t.Errorf("expected 3 log levels, got %d", len(levels))
	}

	// Verify min level is set correctly
	if string(filter.MinLevel) != LogLevelDebug {
		t.Errorf("expected min level to be DEBUG, got %s", filter.MinLevel)
	}
}
