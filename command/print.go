package command

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"

	"tmuxist/config"
	"tmuxist/logger"
	"tmuxist/renderer"
)

// PrintCommand represents a print startup script command.
type PrintCommand struct {
	profile string
}

// Name returns the name of PrintCommand.
func (*PrintCommand) Name() string {
	return "print"
}

// Synopsis returns a short string describing PrintCommand.
func (*PrintCommand) Synopsis() string {
	return "print tmuxist configuration"
}

// Usage returns a long string explaining PrintCommand and givinig usage.
func (*PrintCommand) Usage() string {
	return "print: tmuxist print [-profile profile]\n"
}

// SetFlags adds the flags for PrintCommand to the specified set.
func (cmd *PrintCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.profile, "profile", config.DefaultProfileName(), "Profile")
}

// Execute executes print startup script and returns an ExitStatus.
func (cmd *PrintCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	c, err := config.LoadFileByProfile(cmd.profile)
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	r := renderer.StartRenderer{Config: c}
	fmt.Printf("%s", r.Render())

	return subcommands.ExitSuccess
}
