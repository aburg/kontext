package util

import "os"

func DeleteFile(path string) error {
	return os.Remove(path)
}
