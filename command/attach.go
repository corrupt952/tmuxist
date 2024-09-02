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

// AttachCommand represents a attach tmux session command.
type AttachCommand struct {}

// Name returns the name of AttachCommand.
func (*AttachCommand) Name() string {
	return "attach"
}

// Synopsis returns a short string describing AttachCommand.
func (*AttachCommand) Synopsis() string {
	return "attach tmux session"
}

// Usage returns a long string explaining AttachCommand and giving usage.
func (*AttachCommand) Usage() string {
	return "kill: tmuxist kill\n"
}

// SetFlags adds the flags for AttachCommand to the specified set.
func (cmd *AttachCommand) SetFlags(f *flag.FlagSet) {}

// Execute executes attach tmux session and returns an ExitStatus.
func (cmd *AttachCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
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

	r := renderer.AttachRenderer{Config: c}
	if err := shell_helper.Exec(r.Render()); err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
