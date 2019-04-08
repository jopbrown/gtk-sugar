package fs

import (
	"os"

	"fmt"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func CheckExists(filename string) error {
	if Exists(filename) {
		return nil
	}

	return fmt.Errorf("file or dir not exist: %s", filename)
}

func ExistsFile(filename string) bool {
	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return stat.Mode().IsRegular()
}

func CheckExistsFile(filename string) error {
	if ExistsFile(filename) {
		return nil
	}

	return fmt.Errorf("file not exist: %s", filename)
}

func ExistsDir(filename string) bool {
	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return stat.IsDir()
}

func CheckExistsDir(filename string) error {
	if ExistsDir(filename) {
		return nil
	}

	return fmt.Errorf("dir not exist: %s", filename)
}
