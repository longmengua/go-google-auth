package util

import (
	"os"
)

func GetRootPath() (string, error) {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return wd, nil
}
