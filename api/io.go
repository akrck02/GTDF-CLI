package api

import "os"

func SecureMkdir(path string, perm os.FileMode) {
	err := os.MkdirAll(path, perm)
	if err != nil {
		Error(err.Error())
	}
}

func SecureRemoveFile(path string) {
	err := os.Remove(path)
	if err != nil {
		Error(err.Error())
	}
}

func SecureRemoveAll(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		Error(err.Error())
	}
}

func SecureRenameFile(oldPath string, newPath string) {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		Error(err.Error())
	}
}

func RemoveAnyway(path string) {
	os.RemoveAll(path)
}
