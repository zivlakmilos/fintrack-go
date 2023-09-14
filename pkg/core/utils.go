package core

import (
	"os"
	"path"
)

func getBaseDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	baseDir := path.Join(homeDir, ".fintrack")

	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		err := os.Mkdir(baseDir, 0755)
		if err != nil {
			return "", err
		}
	}

	return baseDir, nil
}
