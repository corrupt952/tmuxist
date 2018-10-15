package path

import (
	"os/user"
	"strings"
)

func AbsolutePath(path string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return strings.Replace(path, "~", usr.HomeDir, 1), nil
}
