package command

import (
	"bytes"
	"context"
	"flag"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"text/template"

	"github.com/google/subcommands"

	"tmuxist/config"
	path_helper "tmuxist/helper/path"
	"tmuxist/logger"
)

// InitCommand represents a create configuration command.
type InitCommand struct {
	profile string
}

// Name returns the name of InitCommand.
func (*InitCommand) Name() string {
	return "init"
}

// Synopsis returns a short string describing InitCommand.
func (*InitCommand) Synopsis() string {
	return "initialize tmuxist configuration"
}

// Usage returns a long string explaining InitCommand and givinig usage.
func (*InitCommand) Usage() string {
	return "init: tmuxist init [-profile profile]\n"
}

// SetFlags adds the flags for InitCommand to the specified set.
func (cmd *InitCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.profile, "profile", "default", "Profile")
}

// Execute executes create configuration and returns an ExitStatus.
func (cmd *InitCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	cfgDirPath, err := path_helper.Fullpath(config.ConfigDirPath)
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}
	if err := os.MkdirAll(cfgDirPath, os.ModePerm); err != nil {
		logger.Warn(err.Error())
	}

	cfgPath := filepath.Join(cfgDirPath, cmd.profile+".toml")
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
	usr, err := user.Current()
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}
	var buf bytes.Buffer
	attach := true
	err = tmpl.Execute(&buf, &config.Config{
		Name:    cmd.profile,
		Root:    usr.HomeDir,
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
