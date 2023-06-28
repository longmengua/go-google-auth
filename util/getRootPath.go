package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetRootPath() (string, error) {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Traverse upwards in the directory hierarchy until we find a directory
	// that contains a known file (e.g., main.go or go.mod) indicating the root of the project
	for {
		files, err := filepath.Glob(filepath.Join(wd, "*.go")) // Adjust the pattern based on your project's files
		if err != nil {
			return "", err
		}

		if len(files) > 0 {
			return wd, nil
		}

		// Move up one directory level
		wd = filepath.Dir(wd)
		if wd == "/" || wd == "." {
			break
		}
	}

	return "", fmt.Errorf("root path not found")
}
