package renderer

import (
	"fmt"
	"regexp"
	"strconv"
)

// parseGridLayout parses grid layout notation (e.g. "2x2", "3x2") and returns
// the number of columns and rows, or an error if the format is invalid
func parseGridLayout(layout string) (cols, rows int, err error) {
	// Match patterns like "2x2", "3x4", etc.
	re := regexp.MustCompile(`^(\d+)x(\d+)$`)
	matches := re.FindStringSubmatch(layout)

	if len(matches) != 3 {
		return 0, 0, fmt.Errorf("invalid grid layout format: %s", layout)
	}

	cols, err = strconv.Atoi(matches[1])
	if err != nil {
		return 0, 0, err
	}

	rows, err = strconv.Atoi(matches[2])
	if err != nil {
		return 0, 0, err
	}

	return cols, rows, nil
}

// isGridLayout checks if the layout string is a grid notation
func isGridLayout(layout string) bool {
	re := regexp.MustCompile(`^\d+x\d+$`)
	return re.MatchString(layout)
}

// generateGridCommands generates tmux commands to create a grid layout
func generateGridCommands(cols, rows int) string {
	if cols <= 0 || rows <= 0 {
		return ""
	}

	// For a grid layout, we'll use tmux's tiled layout as a base
	// and then adjust if needed
	// totalPanes := cols * rows // Reserved for future use

	// Simple approach: use tiled layout for now
	// This will automatically arrange panes in a grid-like pattern
	return "tmux select-layout tiled\n"
}
