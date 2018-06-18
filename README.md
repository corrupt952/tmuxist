# tmuxist

[![CircleCI](https://circleci.com/gh/corrupt952/tmuxist.svg?style=svg)](https://circleci.com/gh/corrupt952/tmuxist)

## TODO
### main
* layout
* list profiles
* attach profile
* log level
### circleci
* build automation
### test
* config
* logger

## Commands
### tmuxist init
Initialize configuration.

```sh
tmuxist init
# or, with profile name
tmuxist init -profile your_profile
```

### tmuxist edit
Edit configuration.

```sh
tmuxist edit
# or, with profile name
tmuxist edit -profile your_profile
```

### tmuxist start
Start tmux session.

```sh
tmuxist start
# or, with profile name
tmuxist start -profile your_profile
```


## Configuration

```toml
name    = 'tmuxist'
root    = '~'
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
[[windows.panes]]
[[windows.panes]]
```
