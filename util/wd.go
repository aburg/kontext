package util

import "os"

func getWd() string {
	pw, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return pw
}
