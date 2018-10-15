package command

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"

	"github.com/corrupt952/tmuxist/config"
	"github.com/corrupt952/tmuxist/logger"
	"github.com/corrupt952/tmuxist/renderer"
)

type PrintCommand struct {
	profile string
}

func (*PrintCommand) Name() string {
	return "print"
}
func (*PrintCommand) Synopsis() string {
	return "print tmuxist configuration"
}
func (*PrintCommand) Usage() string {
	return "print: tmuxist print [-profile profile]\n"
}
func (cmd *PrintCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.profile, "profile", "default", "Profile")
}
func (cmd *PrintCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	c, err := config.LoadFileByProfile(cmd.profile)
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	r := renderer.StartRenderer{c}
	fmt.Printf("%s", r.Render())

	return subcommands.ExitSuccess
}
