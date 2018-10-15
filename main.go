package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"

	"github.com/corrupt952/tmuxist/command"
	"github.com/corrupt952/tmuxist/logger"
)

func main() {
	logger.Setup(os.Stderr)

	subcommands.Register(&command.InitCommand{}, "")
	subcommands.Register(&command.EditCommand{}, "")
	subcommands.Register(&command.PrintCommand{}, "")
	subcommands.Register(&command.StartCommand{}, "")
	subcommands.Register(&command.VersionCommand{}, "")
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
