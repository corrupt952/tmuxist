package command

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"tmuxist/config"
	"tmuxist/logger"

	"github.com/google/subcommands"
)

// LIstCommand represents a version command.
type LIstCommand struct{}

// Name returns the name of LIstCommand.
func (*LIstCommand) Name() string {
	return "list"
}

// Synopsis returns a short string describing LIstCommand.
func (*LIstCommand) Synopsis() string {
	return "List tmuxist profiles"
}

// Usage returns a long string explaining LIstCommand and givinig usage.
func (*LIstCommand) Usage() string {
	return "list: show tmuxist profiles\n"
}

// SetFlags adds the flags for LIstCommand to the specified set.
func (*LIstCommand) SetFlags(f *flag.FlagSet) {
}

// Execute executes print version and returns an ExitStatus.
func (*LIstCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	path, err := config.ConfigurationDirectoryPath()
	if err != nil {
		logger.Err(err.Error())
		logger.Err("Please execute: `tmuxist init`")
		return subcommands.ExitFailure
	}

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		_, err = filepath.Match("*.toml", path)
		if err != nil {
			return err
		}

		c, err := config.LoadFile(path)
		if err != nil {
			return err
		}
		fmt.Println(c.Name)
		return nil
	})
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
