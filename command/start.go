package command

import (
	"context"
	"flag"
	"os"
	"strings"

	"github.com/google/subcommands"

	"tmuxist/config"
	shell_helper "tmuxist/helper/shell"
	"tmuxist/logger"
	renderer "tmuxist/renderer"
)

// StartCommand represents a startup tmux session command.
type StartCommand struct {}

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
func (cmd *StartCommand) SetFlags(f *flag.FlagSet) {}

// HasSession checks if the session is already exists.
func (cmd *StartCommand) HasSession(c *config.Config) bool {
	r := renderer.ListSessionsRenderer{}
	output, err := shell_helper.ExecWithOutput(r.Render())
	if err != nil {
		logger.Err(err.Error())
		return false
	}
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if line == c.Name {
			return true
		}
	}
	return false
}

// Execute executes startup tmux session and returns an ExitStatus.
func (cmd *StartCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
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

	if cmd.HasSession(c) {
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
