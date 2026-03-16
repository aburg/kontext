package util

import (
	"fmt"
	"os"
)

func SetFileContent(kind string, value string) error {
	filename := CreateKontextFilename(kind)

	// (using pwd)
	filepath := filename

	exists, err := FileExists(filepath)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("that context is already set at %s", filepath)
	}

	err = os.WriteFile(filename, []byte(value), 0o666)
	if err != err {
		return err
	}

	return nil
}
