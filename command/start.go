package command

import (
	"context"
	"flag"

	"github.com/google/subcommands"

	shell_helper "tmuxist/helper/shell"
	"tmuxist/logger"
	renderer "tmuxist/renderer"
	"tmuxist/session"
)

// StartCommand represents a startup tmux session command.
type StartCommand struct {
	ConfigCommand
}

// Name returns the name of StartCommand.
func (*StartCommand) Name() string {
	return "start"
}

// Synopsis returns a short string describing StartCommand.
func (*StartCommand) Synopsis() string {
	return "start tmux session"
}

// Usage returns a long string explaining StartCommand and giving usage.
func (*StartCommand) Usage() string {
	return "start: tmuxist start\n"
}

// SetFlags adds the flags for StartCommand to the specified set.
func (cmd *StartCommand) SetFlags(f *flag.FlagSet) {
	cmd.SetConfigFlags(f)
}

// Execute executes startup tmux session and returns an ExitStatus.
func (cmd *StartCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	c, err := cmd.LoadConfig()
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	hasSession, err := session.HasSession(c.Name)
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	if hasSession {
		r := renderer.AttachRenderer{Config: c}
		if err := shell_helper.Exec(r.Render()); err != nil {
			logger.Err(err.Error())
			return subcommands.ExitFailure
		}
	} else {
		r := renderer.StartRenderer{Config: c}
		if err := shell_helper.Exec(r.Render()); err != nil {
			logger.Err(err.Error())
			return subcommands.ExitFailure
		}
	}

	return subcommands.ExitSuccess
}
