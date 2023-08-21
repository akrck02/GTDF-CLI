package core

import (
	"errors"
	"os"
)

func IsProject(directory string) bool {

	file, err := os.Open(directory + "/gtdf.config.json")
	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	file.Close()
	return true
}
