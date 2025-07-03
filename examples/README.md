# tmuxist Examples

This directory contains various example configurations for tmuxist, demonstrating all available features and common use cases.

## Basic Examples

- **basic.yaml** - Simplest configuration with a single window and pane
- **multiple-windows.yaml** - Creating multiple windows in a session
- **standard-layouts.yaml** - All standard tmux layout options (even-horizontal, main-vertical, etc.)

## AI-Assisted Development

- **claude-code.yaml** - Simple Claude Code setup (similar to this repo's .tmuxist.yaml)
- **ai-assisted-dev.yaml** - Combining Claude Code with traditional dev tools
- **claude-code-advanced.yaml** - Advanced setup with specialized Claude instances for different contexts

## New Features (v1.2.0)

- **grid-layouts.yaml** - Grid notation for easy pane arrangement (e.g., "2x2", "3x2")
- **pane-sizes.yaml** - Percentage-based pane sizing (e.g., size: "30%")

## Advanced Examples

- **combined-features.yaml** - Using multiple features together
- **sync-panes.yaml** - Synchronized panes for running commands simultaneously
- **real-world-webapp.yaml** - Practical full-stack development setup
- **edge-cases.yaml** - Edge cases and special scenarios for testing

## Usage

To use any of these examples:

```bash
# Start a session using an example configuration
tmuxist start -f examples/basic.yaml

# Start without attaching
tmuxist start -f examples/grid-layouts.yaml --no-attach

# Kill a session
tmuxist kill -f examples/basic.yaml
```

## Features Demonstrated

### Grid Layouts
- `"2x2"` - 2 columns × 2 rows (4 panes)
- `"3x2"` - 3 columns × 2 rows (6 panes)
- `"4x1"` - 4 columns × 1 row (4 panes horizontally)
- `"1x4"` - 1 column × 4 rows (4 panes vertically)

### Pane Sizes
- `size: "30%"` - Pane takes 30% of remaining space
- Sizes are relative to the remaining space after previous splits
- Only applies to non-first panes in a window

### Standard Layouts
- `even-horizontal` - Evenly split horizontally
- `even-vertical` - Evenly split vertically
- `main-horizontal` - Large pane on top
- `main-vertical` - Large pane on left
- `tiled` - Arrange panes in a grid

### Other Features
- `sync: true` - Synchronize input across all panes in a window
- Multi-line commands using YAML's `|` syntax
- Setting custom session names and root directories