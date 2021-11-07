package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"

	"tmuxist/command"
	"tmuxist/logger"
)

func main() {
	logger.Setup(os.Stderr)

	subcommands.Register(&command.LIstCommand{}, "")
	subcommands.Register(&command.InitCommand{}, "")
	subcommands.Register(&command.EditCommand{}, "")
	subcommands.Register(&command.PrintCommand{}, "")
	subcommands.Register(&command.StartCommand{}, "")
	subcommands.Register(&command.KillCommand{}, "")
	subcommands.Register(&command.AttachCommand{}, "")
	subcommands.Register(&command.VersionCommand{}, "")
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
