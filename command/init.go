package command

import (
	"bytes"
	"context"
	"flag"
	"io/ioutil"
	"os"
	"text/template"
	"path/filepath"

	"github.com/google/subcommands"

	"tmuxist/config"
	"tmuxist/logger"
)

// InitCommand represents a create configuration command.
type InitCommand struct{}

// Name returns the name of InitCommand.
func (*InitCommand) Name() string {
	return "init"
}

// Synopsis returns a short string describing InitCommand.
func (*InitCommand) Synopsis() string {
	return "initialize tmuxist configuration"
}

// Usage returns a long string explaining InitCommand and giving usage.
func (*InitCommand) Usage() string {
	return "init: tmuxist init\n"
}

// SetFlags adds the flags for InitCommand to the specified set.
func (cmd *InitCommand) SetFlags(f *flag.FlagSet) {}

// Execute executes create configuration and returns an ExitStatus.
func (cmd *InitCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	cfgPath, err := config.ConfigurationPath()
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	if _, err := os.Stat(cfgPath); err == nil {
		logger.Warn(cfgPath + " is already exists.")
		return subcommands.ExitFailure
	}
	tmpl, err := template.New("tmuxist").Parse(`name = "{{.Name}}"
root = "{{.Root}}"
attach = {{.Attach}}

[[windows]]
[[windows.panes]]
command = "echo 'hello'"`)
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}
	currentPath, err := os.Getwd()
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}
	directory := filepath.Base(currentPath)
	var buf bytes.Buffer
	attach := true
	err = tmpl.Execute(&buf, &config.Config{
		Name:    directory,
		Root:    ".",
		Attach:  &attach,
		Windows: []config.Window{},
	})
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}
	err = ioutil.WriteFile(cfgPath, buf.Bytes(), 0644)
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
