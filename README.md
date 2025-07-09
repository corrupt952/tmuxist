# tmuxist

[![Test](https://github.com/corrupt952/tmuxist/actions/workflows/test.yaml/badge.svg)](https://github.com/corrupt952/tmuxist/actions/workflows/test.yaml)

`tmuxist` is a tool to manage tmux sessions with configuration file.
You can define tmux session in `.tmuxist.yaml` (or `.tmuxist.yml`, `tmuxist.toml`) and start or attach tmux session with `tmuxist` command.

## Installation

### Manual

Download tar.gz on [Latest release](https://github.com/corrupt952/tmuxist/releases/latest).

### Homebrew

```sh
brew tap corrupt952/tmuxist
brew install tmuxist
```

## Commands

### tmuxist init

Initialize configuration.
When you run this command, `.tmuxist.yaml` is created in the current directory.

```sh
tmuxist init
```

You can also specify the format:

```sh
tmuxist init --format toml  # Creates tmuxist.toml
tmuxist init --format yaml  # Creates .tmuxist.yaml (default)
```

### tmuxist start

Start or attach tmux session.
When you run this command, the session defined in `.tmuxist.yaml` is created or attached.

```sh
tmuxist start
```

You can also specify a custom configuration file:

```sh
tmuxist start -f custom-config.yaml
```

### tmuxist kill

Kill tmux session.
When you run this command, the session defined in `.tmuxist.yaml` is deleted.

```sh
tmuxist kill
```

You can also specify a custom configuration file:

```sh
tmuxist kill -f custom-config.yaml
```

## Layouts

tmuxist supports all standard tmux layouts plus a new grid notation for easy pane arrangement.

### Standard Layouts

#### even-horizontal

Panes are spread out evenly from left to right.

```
┌─────────┬─────────┬─────────┐
│         │         │         │
│  Pane1  │  Pane2  │  Pane3  │
│         │         │         │
└─────────┴─────────┴─────────┘
```

#### even-vertical

Panes are spread out evenly from top to bottom.

```
┌─────────────────┐
│     Pane1       │
├─────────────────┤
│     Pane2       │
├─────────────────┤
│     Pane3       │
└─────────────────┘
```

#### main-horizontal

One large pane on top, others spread out evenly below.

```
┌─────────────────┐
│                 │
│   Main Pane     │
│                 │
├────────┬────────┤
│ Pane2  │ Pane3  │
└────────┴────────┘
```

#### main-vertical

One large pane on the left, others spread out evenly on the right.

```
┌─────────┬───────┐
│         │ Pane2 │
│  Main   ├───────┤
│  Pane   │ Pane3 │
│         ├───────┤
│         │ Pane4 │
└─────────┴───────┘
```

#### tiled

Panes are spread out as evenly as possible in both rows and columns.

```
┌────────┬────────┐
│ Pane1  │ Pane2  │
├────────┼────────┤
│ Pane3  │ Pane4  │
└────────┴────────┘
```

### Grid Notation (New Feature)

Use simple grid notation like "2x2", "3x2" for easy pane arrangement.

#### "2x2" - 2 columns × 2 rows

```
┌────────┬────────┐
│ Pane1  │ Pane2  │
├────────┼────────┤
│ Pane3  │ Pane4  │
└────────┴────────┘
```

#### "3x2" - 3 columns × 2 rows

```
┌──────┬──────┬──────┐
│ Pane1│ Pane2│ Pane3│
├──────┼──────┼──────┤
│ Pane4│ Pane5│ Pane6│
└──────┴──────┴──────┘
```

#### "1x4" - 1 column × 4 rows

```
┌─────────────────┐
│     Pane1       │
├─────────────────┤
│     Pane2       │
├─────────────────┤
│     Pane3       │
├─────────────────┤
│     Pane4       │
└─────────────────┘
```

### Example Usage

```yaml
# Using standard layout
windows:
  - layout: main-vertical
    panes:
      - command: vim
      - command: npm run dev
      - command: npm test

# Using grid notation
windows:
  - layout: "2x2"
    panes:
      - command: htop
      - command: docker stats
      - command: tail -f app.log
      - command: watch date
```

## Architecture

`tmuxist` reads `.tmuxist.yaml` (or `.tmuxist.yml`, `tmuxist.toml`) in the directory where the command is executed and manages tmux sessions.

### Configuration

Configuration files can be written in YAML or TOML format. YAML is the default and recommended format.

#### YAML Configuration Example

In the example of `.tmuxist.yaml` below, the following tmux session is created.

- Session name ... `tmuxist`
- Window 1
  - Pane 1 ... `htop` command is executed in root(current directory)
- Window 2
  - Pane 1 ... `cd ~/Repo/corrupt952/tmuxist` command is executed and move to the directory
  - Pane 2 ... Empty pane
- Window 3
  - Layout ... `tiled`
  - Synchronize panes ... `true`
  - Pane 1 ... Empty pane
  - Pane 2 ... Empty pane
  - Pane 3 ... Empty pane
  - Pane 4 ... Empty pane

```yaml
name: tmuxist
root: .
attach: true

windows:
  # Window 1
  - name: "Monitor"
    panes:
      - command: htop

  # Window 2
  - name: "Development"
    panes:
      - command: cd ~/Repo/corrupt952/tmuxist
      - command: ""  # Empty pane

  # Window 3
  - name: "Servers"
    layout: tiled
    sync: true
    panes:
      - command: ""  # Empty pane
      - command: ""  # Empty pane
      - command: ""  # Empty pane
      - command: ""  # Empty pane
```

#### Environment Variables (tmux 3.0+)

You can set environment variables for the entire session using the `env` field at the root level:

```yaml
name: tmuxist
root: .
attach: true
env:
  NODE_ENV: development
  PORT: "3000"
  DEBUG: "app:*"

windows:
  - panes:
      - command: npm run dev
      - command: npm test
```

#### TOML Configuration Example

If you prefer TOML format, create `tmuxist.toml`:

```toml
name    = 'tmuxist'
root    = '.'
attach  = true

[[windows]]
name = "Monitor"
[[windows.panes]]
command = "htop"

[[windows]]
name = "Development"
[[windows.panes]]
command = "cd ~/Repo/corrupt952/tmuxist"
[[windows.panes]]
command = ""

[[windows]]
name = "Servers"
layout = "tiled"
sync = true
[[windows.panes]]
command = ""
[[windows.panes]]
command = ""
[[windows.panes]]
command = ""
[[windows.panes]]
command = ""
```

Environment variables in TOML:

```toml
name = "tmuxist"
root = "."
attach = true

[env]
NODE_ENV = "development"
PORT = "3000"
DEBUG = "app:*"

[[windows]]
[[windows.panes]]
command = "npm run dev"
```
