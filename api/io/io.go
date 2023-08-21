package io

import (
	"os"

	"github.com/akrck02/GTDF-CLI/api/logger"
)

func SecureMkdir(path string, perm os.FileMode) {
	err := os.MkdirAll(path, perm)
	if err != nil {
		logger.Error(err.Error())
	}
}

func SecureRemoveFile(path string) {
	err := os.Remove(path)
	if err != nil {
		logger.Error(err.Error())
	}
}

func SecureRemoveAll(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		logger.Error(err.Error())
	}
}

func SecureRenameFile(oldPath string, newPath string) {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		logger.Error(err.Error())
	}
}

func RemoveAnyway(path string) {
	os.RemoveAll(path)
}

func GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		logger.Error(err.Error())
	}
	return dir
}

func WriteToFile(file string, text string) {

	// add this line to existing file if it exists,
	// create it if it doesn't

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error(err.Error())
	}
	defer f.Close()

	if _, err := f.WriteString(text); err != nil {
		logger.Error(err.Error())
	}

}

func WriteFileWith(file string, text string) {

	// add this line to existing file if it exists,
	// create it if it doesn't

	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error(err.Error())
	}
	defer f.Close()

	if _, err := f.WriteString(text); err != nil {
		logger.Error(err.Error())
	}

}
