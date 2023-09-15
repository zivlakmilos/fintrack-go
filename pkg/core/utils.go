package core

import (
	"fmt"
	"os"
	"path"
)

func GetBaseDir() (string, error) {
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

func GetDBPath(year int) (string, error) {
	baseDir, err := GetBaseDir()
	if err != nil {
		return "", err
	}

	dbPath := path.Join(baseDir, fmt.Sprintf("%d.db", year))

	return dbPath, nil
}
