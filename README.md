# tmuxist

[![Test](https://github.com/corrupt952/tmuxist/actions/workflows/test.yaml/badge.svg)](https://github.com/corrupt952/tmuxist/actions/workflows/test.yaml)

`tmuxist` is a tool to manage tmux sessions with configuration file.  
You can define tmux session in `tmuxist.toml` and start or attach tmux session with `tmuxist` command.

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
When you run this command, `tmuxist.toml` is created in the current directory.

```sh
tmuxist init
```

### tmuxist start

Start or attach tmux session.  
When you run this command, the session defined in `tmuxist.toml` is created or attached.

```sh
tmuxist start
```

### tmuxist kill

Kill tmux session.  
When you run this command, the session defined in `tmuxist.toml` is deleted.

```sh
tmuxist kill
```

## Architecture

`tmuxist` reads `tmuxist.toml` in the directory where the command is executed and manages tmux sessions.  

### Configuration

`tmuxist.toml` is a configuration file written in TOML format.

In the example of `tmuxist.toml` below, the following tmux session is created.

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

```toml
name    = 'tmuxist'
root    = '.'
attach  = true

[[windows]]
[[windows.panes]]
command = """
htop
"""

[[windows]]
[[windows.panes]]
command = """
cd ~/Repo/corrupt952/tmuxist
"""
[[windows.panes]]

[[windows]]
layout = "tiled"
sync = true
[[windows.panes]]
[[windows.panes]]
[[windows.panes]]
[[windows.panes]]
```
