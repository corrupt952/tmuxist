package shell

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

// CurrentShell returns SHELL environment variables.
func CurrentShell() string {
	return os.Getenv("SHELL")
}

// CommandSubstitution returns command converted to command substitution.
func CommandSubstitution(s string) string {
	shell := filepath.Base(CurrentShell())

	switch shell {
	case "bash", "zsh":
		return fmt.Sprintf("$(%s)", s)
	default:
		return fmt.Sprintf("`%s`", s)
	}
}

// Exec executes command on current shell
func Exec(command string) error {
	return syscall.Exec(
		"/bin/sh",
		[]string{
			CurrentShell(),
			"-cx",
			command,
		},
		os.Environ(),
	)
}

// ExecWithOutput executes command on current shell and returns output
func ExecWithOutput(command string) (string, error) {
	cmd := exec.Command(CurrentShell(), "-c", command)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
