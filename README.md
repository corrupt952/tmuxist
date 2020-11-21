# tmuxist

[![CircleCI](https://circleci.com/gh/corrupt952/tmuxist.svg?style=svg)](https://circleci.com/gh/corrupt952/tmuxist)

## TODO
### main
* list profiles
* delete profile
* log level
* ADD TMUXIST_PROFILE environment variable
### test
* config
* logger

## Installation
Download tar.gz on [Latest release](https://github.com/corrupt952/tmuxist/releases/latest).

### Homebrew
```sh
brew tap corrupt952/tmuxist
brew install tmuxist
```

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
# or, specify editor
tmuxist edit -profile gvim
```

### tmuxist start
Start tmux session.

```sh
tmuxist start
# or, with profile name
tmuxist start -profile your_profile
```

### tmuxist output
Ouput tmuxist configuration.

```sh
tmuxist output
# or, with profile name
tmuxist output -profile your_profile
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
layout = "tiled"
sync = true
[[windows.panes]]
[[windows.panes]]
[[windows.panes]]
[[windows.panes]]
```
