package command

import (
	"context"
	"flag"
	"os"
	"syscall"

	"github.com/google/subcommands"

	"github.com/corrupt952/tmuxist/config"
	shell_helper "github.com/corrupt952/tmuxist/helper/shell"
	"github.com/corrupt952/tmuxist/logger"
	"github.com/corrupt952/tmuxist/renderer"
)

type StartCommand struct {
	profile string
}

func (*StartCommand) Name() string {
	return "start"
}
func (*StartCommand) Synopsis() string {
	return "start tmux session"
}
func (*StartCommand) Usage() string {
	return "start: tmuxist start [-profile profile]\n"
}
func (cmd *StartCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.profile, "profile", "default", "Profile")
}
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
	if err := syscall.Exec("/bin/sh", []string{shell_helper.CurrentShell(), "-c", r.Render()}, os.Environ()); err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
