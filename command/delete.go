package command

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"

	"tmuxist/config"
	"tmuxist/logger"
)

// DeleteCommand represents a print startup script command.
type DeleteCommand struct {
	profile string
}

// Name returns the name of DeleteCommand.
func (*DeleteCommand) Name() string {
	return "delete"
}

// Synopsis returns a short string describing DeleteCommand.
func (*DeleteCommand) Synopsis() string {
	return "delete tmuxist configuration"
}

// Usage returns a long string explaining DeleteCommand and giving usage.
func (*DeleteCommand) Usage() string {
	return "delete: tmuxist delete [-profile profile]\n"
}

// SetFlags adds the flags for DeleteCommand to the specified set.
func (cmd *DeleteCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.profile, "profile", config.DefaultProfileName(), "Profile")
}

// Execute executes print startup script and returns an ExitStatus.
func (cmd *DeleteCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	path, err := config.ConfigurationPath(cmd.profile)
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}
	if _, err := os.Stat(path); err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}
	if err := os.Remove(path); err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	fmt.Println("Delete profile: " + cmd.profile)
	return subcommands.ExitSuccess
}
