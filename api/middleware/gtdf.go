package middleware

import (
	"errors"
	"os"

	"github.com/akrck02/GTDF-CLI/api/io"
	"github.com/akrck02/GTDF-CLI/api/logger"
	"github.com/spf13/cobra"
)

func IsProject(directory string) bool {

	file, err := os.Open(directory + "/gtdf.config.json")
	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	file.Close()
	return true
}

func ProjectCheck(cmd *cobra.Command, args []string) {

	logger.Log("Current directory: " + io.GetCurrentDirectory())
	if !IsProject(io.GetCurrentDirectory()) {
		logger.Warn("You are not in a GTDF project directory.")
		os.Exit(1)
	}

}
