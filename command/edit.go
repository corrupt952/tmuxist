package command

import (
	"context"
	"flag"
	"os"
	"os/exec"

	"github.com/google/subcommands"

	"tmuxist/config"
	"tmuxist/logger"
)

// EditCommand represents a edit configuration command.
type EditCommand struct {
	profile string
	editor  string
}

// Name returns the name of EditCommand.
func (*EditCommand) Name() string {
	return "edit"
}

// Synopsis returns a short string describing EditCommand.
func (*EditCommand) Synopsis() string {
	return "edit tmuxist configuration"
}

// Usage returns a long string explaining EditCommand and givinig usage.
func (*EditCommand) Usage() string {
	return "edit: tmuxist edit [-editor editor] [-profile profile]\n"
}

// SetFlags adds the flags for EditCommand to the specified set.
func (cmd *EditCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.profile, "profile", config.DefaultProfileName(), "Profile")
	f.StringVar(&cmd.editor, "editor", "vim", "Editor")
}

// Execute executes edit configuration and returns an ExitStatus.
func (cmd *EditCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	editor := cmd.editor
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}
	if editor == "" {
		logger.Err("$EDITOR or -editor is required")
		return subcommands.ExitFailure
	}

	path, err := config.ConfigurationPath(cmd.profile)
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}
	if _, err := os.Stat(path); err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	shell := exec.Command(editor, path)
	shell.Stdin = os.Stdin
	shell.Stdout = os.Stdout
	shell.Stderr = os.Stderr
	if err = shell.Run(); err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
