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

	"github.com/corrupt952/tmuxist/config"
	path_helper "github.com/corrupt952/tmuxist/helper/path"
	"github.com/corrupt952/tmuxist/logger"
)

type InitCommand struct {
	profile string
}

func (*InitCommand) Name() string {
	return "init"
}
func (*InitCommand) Synopsis() string {
	return "initialize tmuxist configuration"
}
func (*InitCommand) Usage() string {
	return "init: tmuxist init [-profile profile]\n"
}
func (cmd *InitCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.profile, "profile", "default", "Profile")
}
func (cmd *InitCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	cfgDirPath, err := path_helper.AbsolutePath(config.ConfigDirPath)
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
	var attach bool = true
	err = tmpl.Execute(&buf, &config.Config{cmd.profile, usr.HomeDir, &attach, []config.Window{}})
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
