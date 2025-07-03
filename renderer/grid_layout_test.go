package renderer

import (
	"testing"
	test_helper "tmuxist/helper/test"
)

func TestParseGridLayout(t *testing.T) {
	tests := []struct {
		name     string
		layout   string
		wantCols int
		wantRows int
		wantErr  bool
	}{
		{"valid 2x2", "2x2", 2, 2, false},
		{"valid 3x2", "3x2", 3, 2, false},
		{"valid 1x4", "1x4", 1, 4, false},
		{"invalid format", "2by2", 0, 0, true},
		{"invalid format no x", "22", 0, 0, true},
		{"invalid format empty", "", 0, 0, true},
		{"invalid format text", "axb", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cols, rows, err := parseGridLayout(tt.layout)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				test_helper.AssertEquals(t, cols, tt.wantCols)
				test_helper.AssertEquals(t, rows, tt.wantRows)
			}
		})
	}
}

func TestIsGridLayout(t *testing.T) {
	tests := []struct {
		name   string
		layout string
		want   bool
	}{
		{"valid 2x2", "2x2", true},
		{"valid 10x10", "10x10", true},
		{"invalid tiled", "tiled", false},
		{"invalid main-vertical", "main-vertical", false},
		{"invalid 2by2", "2by2", false},
		{"invalid empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isGridLayout(tt.layout)
			test_helper.AssertEquals(t, got, tt.want)
		})
	}
}

func TestGenerateGridCommands(t *testing.T) {
	tests := []struct {
		name string
		cols int
		rows int
		want string
	}{
		{"2x2 grid", 2, 2, "tmux select-layout tiled\n"},
		{"3x3 grid", 3, 3, "tmux select-layout tiled\n"},
		{"invalid 0x0", 0, 0, ""},
		{"invalid negative", -1, 2, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateGridCommands(tt.cols, tt.rows)
			test_helper.AssertEquals(t, got, tt.want)
		})
	}
}
