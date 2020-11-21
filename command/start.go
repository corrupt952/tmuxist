package command

import (
	"context"
	"flag"

	"github.com/google/subcommands"

	"tmuxist/config"
	shell_helper "tmuxist/helper/shell"
	"tmuxist/logger"
	"tmuxist/renderer"
)

// StartCommand represents a startup tmux session command.
type StartCommand struct {
	profile string
}

// Name returns the name of StartCommand.
func (*StartCommand) Name() string {
	return "start"
}

// Synopsis returns a short string describing StartCommand.
func (*StartCommand) Synopsis() string {
	return "start tmux session"
}

// Usage returns a long string explaining StartCommand and givinig usage.
func (*StartCommand) Usage() string {
	return "start: tmuxist start [-profile profile]\n"
}

// SetFlags adds the flags for StartCommand to the specified set.
func (cmd *StartCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.profile, "profile", "default", "Profile")
}

// Execute executes startup tmux session and returns an ExitStatus.
func (cmd *StartCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fpath, err := config.ConfigurationPath(cmd.profile)
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	c, err := config.LoadFile(fpath)
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	r := renderer.StartRenderer{c}
	if err := shell_helper.Exec(r.Render()); err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
