package shell

import (
	"fmt"
	"os"
	"path/filepath"
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
