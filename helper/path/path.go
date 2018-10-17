package path

import (
	"os/user"
	"strings"
)

// Fullpath returns path converted to fullpath.
func Fullpath(path string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return strings.Replace(path, "~", usr.HomeDir, 1), nil
}
