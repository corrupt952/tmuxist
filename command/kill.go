package command

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"

	"tmuxist/config"
	shell_helper "tmuxist/helper/shell"
	"tmuxist/logger"
	"tmuxist/renderer"
)

// KillCommand represents a kill tmux session command.
type KillCommand struct {}

// Name returns the name of KillCommand.
func (*KillCommand) Name() string {
	return "kill"
}

// Synopsis returns a short string describing KillCommand.
func (*KillCommand) Synopsis() string {
	return "kill tmux session"
}

// Usage returns a long string explaining KillCommand and giving usage.
func (*KillCommand) Usage() string {
	return "kill: tmuxist kill\n"
}

// SetFlags adds the flags for KillCommand to the specified set.
func (cmd *KillCommand) SetFlags(f *flag.FlagSet) {}

// Execute executes kill tmux session and returns an ExitStatus.
func (cmd *KillCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	path, err := config.ConfigurationPath()
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}
	if _, err := os.Stat(path); err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	c, err := config.LoadFile(path)
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	r := renderer.KillRenderer{Config: c}
	if err := shell_helper.Exec(r.Render()); err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
