package utils

import (
	"os"
	"path/filepath"
)

//EnsureDir makes new directory if the path doesn't exists.
func EnsureDir(fileName string) error {
	dirName := filepath.Dir(fileName)
	if _, err := os.Stat(dirName); err != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			return merr
		}
	}
	return nil
}
