package util

import "path/filepath"

func JoinAbs(elem ...string) (string, error) {
	abs, err := filepath.Abs(filepath.Join(elem...))
	if err != nil {
		return "", err
	}
	return abs, nil
}
