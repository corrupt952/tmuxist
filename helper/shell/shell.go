package shell

import (
	"fmt"
	"os"
	"path/filepath"
)

func CurrentShell() string {
	return os.Getenv("SHELL")
}

func CommandSubstitution(s string) string {
	shell := filepath.Base(CurrentShell())

	switch shell {
	case "bash", "zsh":
		return fmt.Sprintf("$(%s)", s)
	default:
		return fmt.Sprintf("`%s`", s)
	}
}
