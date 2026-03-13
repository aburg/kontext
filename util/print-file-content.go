package util

import (
	"fmt"
	"os"
	"strings"
)

func PrintFileContent(filepath string) error {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	trimmed := strings.TrimSuffix(string(content), "\n")
	fmt.Println(trimmed)
	return nil
}
