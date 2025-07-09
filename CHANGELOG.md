# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.2.5] - 2025-07-09

### Added

- Window name support using `name` field in window configuration
- Custom window names are set using tmux's rename-window and new-window -n commands
- Example configuration demonstrating window naming feature

## [1.2.4] - 2025-07-09

### Added

- Session-level environment variables support using `env` field in root configuration
- Environment variables are applied to all windows and panes in the session using tmux's `-e` option
- Example configuration for environment variables usage
- Comprehensive example configurations directory with various use cases
- Documentation for layouts with ASCII art visualization

### Changed

- Updated README to reflect YAML as the default configuration format

## [1.2.3] - 2025-07-04

### Changed

- Refactored session management to separate package for better code organization

## [1.2.2] - 2025-07-04

### Changed

- Refactored common config loading logic to base struct for better code reuse

## [1.2.1] - 2025-07-04

### Added

- `-f/--file` flag to specify custom configuration file path

## [1.2.0] - 2025-07-04

### Added

- Grid layout notation (e.g., "2x2", "3x2") for easy pane arrangement
- Percentage-based pane sizing support (e.g., size: "30%")

### Changed

- Migrated default config format from TOML to hidden YAML file (`.tmuxist.yaml`)

## [1.1.0] - 2025-07-03

### Added

- YAML configuration file support (`.tmuxist.yml` and `.tmuxist.yaml`)
- Hidden configuration file support (`.tmuxist.toml`, `.tmuxist.yml`, `.tmuxist.yaml`)
- Made hidden configuration files the default

### Changed

- Updated README documentation
- Updated Go dependencies to v1.24.4

### Fixed

- Updated go-toml/v2 module to v2.2.4

## [1.0.0] - 2024-09-03

### Added

- Initial stable release
- Session management with configuration file
- Window and pane management
- Standard tmux layout support
- Synchronize panes feature
- Commands: `init`, `start`, `kill`