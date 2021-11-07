package command

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

var (
	Version string
)

// VersionCommand represents a version command.
type VersionCommand struct{}

// Name returns the name of VersionCommand.
func (*VersionCommand) Name() string {
	return "version"
}

// Synopsis returns a short string describing VersionCommand.
func (*VersionCommand) Synopsis() string {
	return "Print tmuxist version"
}

// Usage returns a long string explaining VersionCommand and giving usage.
func (*VersionCommand) Usage() string {
	return "version: tmuxist version\n"
}

// SetFlags adds the flags for VersionCommand to the specified set.
func (*VersionCommand) SetFlags(f *flag.FlagSet) {
}

// Execute executes print version and returns an ExitStatus.
func (*VersionCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println(Version)
	return subcommands.ExitSuccess
}
