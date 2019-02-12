package fs

import "os"

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func ExistsFile(filename string) bool {
	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return stat.Mode().IsRegular()
}

func ExistsDir(filename string) bool {
	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return stat.IsDir()
}
