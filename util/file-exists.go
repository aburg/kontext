package util

import (
	"errors"
	"os"
)

func FileExists(file string) (bool, error) {
	info, err := os.Stat(file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		} else {
			return false, err
		}
	}

	if info.IsDir() {
		return false, nil
	} else {
		return true, nil
	}
}
