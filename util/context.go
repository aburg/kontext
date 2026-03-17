package util

import (
	"fmt"
	"os"
)

func FindContext(kind string) (string, error) {
	filename := CreateKontextFilename(kind)

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	testdir := wd
	for {
		testfile, err := JoinAbs(testdir, filename)
		if err != nil {
			return "", err
		}
		exists, err := FileExists(testfile)
		if err != nil {
			return "", err
		}
		if exists {
			return testfile, nil
		} else if testdir == "/" {
			return "", fmt.Errorf("no context file found (looking for \"%s\" from \"%s\" and upwards)", testfile, wd)
		} else {
			testdir, err = JoinAbs(testdir, "..")
			if err != nil {
				return "", fmt.Errorf("could not cd up")
			}
		}
	}
}
