package command

import (
	"context"
	"flag"
	"os"
	"os/exec"

	"github.com/google/subcommands"

	"github.com/corrupt952/tmuxist/config"
	"github.com/corrupt952/tmuxist/logger"
)

type EditCommand struct {
	profile string
	editor  string
}

func (*EditCommand) Name() string {
	return "edit"
}
func (*EditCommand) Synopsis() string {
	return "edit tmuxist configuration"
}
func (*EditCommand) Usage() string {
	return "edit: tmuxist edit [-editor editor] [-profile profile]\n"
}
func (cmd *EditCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.profile, "profile", "default", "Profile")
	f.StringVar(&cmd.editor, "editor", "vim", "Editor")
}
func (cmd *EditCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	editor := cmd.editor
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}
	if editor == "" {
		logger.Err("$EDITOR or -editor is required")
		return subcommands.ExitFailure
	}

	cfgPath, err := config.ConfigurationPath(cmd.profile)
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}
	shell := exec.Command(editor, cfgPath)
	shell.Stdin = os.Stdin
	shell.Stdout = os.Stdout
	shell.Stderr = os.Stderr
	if err = shell.Run(); err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
