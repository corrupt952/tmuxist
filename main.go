package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"
	"text/template"

	"github.com/google/subcommands"
	"github.com/pelletier/go-toml"
)

var (
	version string
)

const (
	CONFIG_DIR_PATH = "~/.config/tmuxist"
	CONFIG_TEMPLATE = `name = "{{.Name}}"
root = "{{.Root}}"
Attach = {{.Attach}}

[[windows]]
[[windows.panes]]
command = "echo 'hello'"`
)

// init commands
type initCmd struct {
	profile string
}

func (*initCmd) Name() string {
	return "init"
}
func (*initCmd) Synopsis() string {
	return "initialize tmuxist configuration"
}
func (*initCmd) Usage() string {
	return "init: tmuxist init [-profile profile]\n"
}
func (cmd *initCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.profile, "profile", "default", "Profile")
}
func (cmd *initCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	cfgDirPath, err := absolutePath(CONFIG_DIR_PATH)
	if err != nil {
		logger.err(err.Error())
		return subcommands.ExitFailure
	}
	if err := os.MkdirAll(cfgDirPath, os.ModePerm); err != nil {
		logger.warn(err.Error())
	}

	cfgPath := filepath.Join(cfgDirPath, cmd.profile+".toml")
	if err != nil {
		logger.err(err.Error())
		return subcommands.ExitFailure
	}

	if _, err := os.Stat(cfgPath); err == nil {
		logger.warn(cfgPath + " is already exists.")
		return subcommands.ExitFailure
	}
	tmpl, err := template.New("tmuxist").Parse(CONFIG_TEMPLATE)
	if err != nil {
		logger.err(err.Error())
		return subcommands.ExitFailure
	}
	usr, err := user.Current()
	if err != nil {
		logger.err(err.Error())
		return subcommands.ExitFailure
	}
	var buf bytes.Buffer
	var attach bool = true
	err = tmpl.Execute(&buf, &Config{cmd.profile, usr.HomeDir, &attach, []Window{}})
	if err != nil {
		logger.err(err.Error())
		return subcommands.ExitFailure
	}
	err = ioutil.WriteFile(cfgPath, buf.Bytes(), 0644)
	if err != nil {
		logger.err(err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

// edit commands
type editCmd struct {
	profile string
	editor  string
}

func (*editCmd) Name() string {
	return "edit"
}
func (*editCmd) Synopsis() string {
	return "edit tmuxist configuration"
}
func (*editCmd) Usage() string {
	return "edit: tmuxist edit [-editor editor] [-profile profile]\n"
}
func (cmd *editCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.profile, "profile", "default", "Profile")
	f.StringVar(&cmd.editor, "editor", "vim", "Editor")
}
func (cmd *editCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	editor := cmd.editor
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}
	if editor == "" {
		logger.err("$EDITOR or -editor is required")
		return subcommands.ExitFailure
	}

	cfgPath, err := getConfigurationPath(cmd.profile)
	if err != nil {
		logger.err(err.Error())
		return subcommands.ExitFailure
	}
	shell := exec.Command(editor, cfgPath)
	shell.Stdin = os.Stdin
	shell.Stdout = os.Stdout
	shell.Stderr = os.Stderr
	if err = shell.Run(); err != nil {
		logger.err(err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

// output commands
type outputCmd struct {
	profile string
}

func (*outputCmd) Name() string {
	return "output"
}
func (*outputCmd) Synopsis() string {
	return "output tmuxist configuration"
}
func (*outputCmd) Usage() string {
	return "output: tmuxist output [-profile profile]\n"
}
func (cmd *outputCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.profile, "profile", "default", "Profile")
}
func (cmd *outputCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	cfgPath, err := getConfigurationPath(cmd.profile)
	if err != nil {
		logger.err(err.Error())
		return subcommands.ExitFailure
	}

	content, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		logger.err(err.Error())
		return subcommands.ExitFailure
	}

	fmt.Printf("%s", content)

	return subcommands.ExitSuccess
}

// start commands
type startCmd struct {
	profile string
}

func (*startCmd) Name() string {
	return "start"
}
func (*startCmd) Synopsis() string {
	return "start tmux session"
}
func (*startCmd) Usage() string {
	return "start: tmuxist start [-profile profile]\n"
}
func (cmd *startCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.profile, "profile", "default", "Profile")
}
func (cmd *startCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fpath, err := getConfigurationPath(cmd.profile)
	if err != nil {
		logger.err(err.Error())
		return subcommands.ExitFailure
	}

	c, err := loadConfiguration(fpath)
	_, err = loadConfiguration(fpath)
	if err != nil {
		logger.err(err.Error())
		return subcommands.ExitFailure
	}

	shell := os.Getenv("SHELL")
	if err := syscall.Exec("/bin/sh", []string{shell, "-c", c.ToScript()}, os.Environ()); err != nil {
		logger.err(err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

// version commands
type versionCmd struct{}

func (*versionCmd) Name() string {
	return "version"
}
func (*versionCmd) Synopsis() string {
	return "Print tmuxist version"
}
func (*versionCmd) Usage() string {
	return "version: tmuxist version\n"
}
func (*versionCmd) SetFlags(f *flag.FlagSet) {
}
func (*versionCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Print(version)
	return subcommands.ExitSuccess
}

func main() {
	initLogger()

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&initCmd{}, "")
	subcommands.Register(&editCmd{}, "")
	subcommands.Register(&outputCmd{}, "")
	subcommands.Register(&startCmd{}, "")
	subcommands.Register(&versionCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}

func getConfigurationPath(profile string) (string, error) {
	fpath, err := absolutePath(filepath.Join(CONFIG_DIR_PATH, profile+".toml"))
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(fpath); err != nil {
		return "", err
	}

	return fpath, nil
}

func loadConfiguration(fpath string) (*Config, error) {
	cfgFile, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	toml.Unmarshal([]byte(cfgFile), &cfg)
	return &cfg, nil
}

func absolutePath(path string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return strings.Replace(path, "~", usr.HomeDir, 1), nil
}
