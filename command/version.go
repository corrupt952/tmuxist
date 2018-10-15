package command

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"

	"github.com/corrupt952/tmuxist/version"
)

type VersionCommand struct{}

func (*VersionCommand) Name() string {
	return "version"
}
func (*VersionCommand) Synopsis() string {
	return "Print tmuxist version"
}
func (*VersionCommand) Usage() string {
	return "version: tmuxist version\n"
}
func (*VersionCommand) SetFlags(f *flag.FlagSet) {
}
func (*VersionCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Print(version.Version)
	return subcommands.ExitSuccess
}
