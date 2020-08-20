package utils

import (
	"os"
)

//EnsureDir ensures if dir is present.
func EnsureDir(dir string) bool {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
