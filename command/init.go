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
type InitCommand struct{
	format string
}

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
func (cmd *InitCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.format, "format", "yaml", "Configuration format (toml, yaml, yml)")
}

// Execute executes create configuration and returns an ExitStatus.
func (cmd *InitCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	// Determine the configuration file path based on format
	currentPath, err := os.Getwd()
	if err != nil {
		logger.Err(err.Error())
		return subcommands.ExitFailure
	}
	
	var cfgPath string
	var configContent string
	
	switch cmd.format {
	case "yaml", "yml":
		cfgPath = filepath.Join(currentPath, "tmuxist.yaml")
		configContent = `name: "{{.Name}}"
root: "{{.Root}}"
attach: {{.Attach}}

windows:
  - panes:
      - command: "echo 'hello'"
`
	default:
		cfgPath = filepath.Join(currentPath, "tmuxist.toml")
		configContent = `name = "{{.Name}}"
root = "{{.Root}}"
attach = {{.Attach}}

[[windows]]
[[windows.panes]]
command = "echo 'hello'"
`
	}

	if _, err := os.Stat(cfgPath); err == nil {
		logger.Warn(cfgPath + " already exists.")
		return subcommands.ExitFailure
	}
	
	tmpl, err := template.New("tmuxist").Parse(configContent)
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
