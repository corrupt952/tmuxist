package session

import (
	"errors"
	"testing"
)

// MockSessionLister is a mock implementation of SessionLister for testing
type MockSessionLister struct {
	sessions []string
	err      error
}

func (m *MockSessionLister) ListSessions() ([]string, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.sessions, nil
}

func TestManager_HasSession(t *testing.T) {
	tests := []struct {
		name         string
		sessions     []string
		sessionName  string
		expectedBool bool
		expectedErr  error
	}{
		{
			name:         "session exists",
			sessions:     []string{"myproject", "another-project", "test-session"},
			sessionName:  "myproject",
			expectedBool: true,
			expectedErr:  nil,
		},
		{
			name:         "session does not exist",
			sessions:     []string{"project1", "project2"},
			sessionName:  "myproject",
			expectedBool: false,
			expectedErr:  nil,
		},
		{
			name:         "empty session list",
			sessions:     []string{},
			sessionName:  "myproject",
			expectedBool: false,
			expectedErr:  nil,
		},
		{
			name:         "error listing sessions",
			sessions:     nil,
			sessionName:  "myproject",
			expectedBool: false,
			expectedErr:  errors.New("failed to list sessions"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockSessionLister{
				sessions: tt.sessions,
				err:      tt.expectedErr,
			}

			manager := NewManager(mock)
			hasSession, err := manager.HasSession(tt.sessionName)

			if tt.expectedErr != nil {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if hasSession != tt.expectedBool {
					t.Errorf("expected HasSession to return %v, but got %v", tt.expectedBool, hasSession)
				}
			}
		})
	}
}

func TestNewManager(t *testing.T) {
	// Test with nil lister (should use default)
	manager := NewManager(nil)
	if manager.lister == nil {
		t.Error("expected manager to have a default lister when nil is passed")
	}

	// Test with custom lister
	mock := &MockSessionLister{}
	manager = NewManager(mock)
	if manager.lister != mock {
		t.Error("expected manager to use the provided lister")
	}
}

func TestDefaultSessionLister_ListSessions(t *testing.T) {
	// This test would require a running tmux server
	// For unit testing, we've tested the logic via mocks above
	// Integration tests would test the actual DefaultSessionLister
	t.Skip("Skipping integration test that requires tmux")
}
