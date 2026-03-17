package util

import (
	"os"

	"github.com/spf13/cobra"
)

func ChangeWorkdir(cmd *cobra.Command) {
	// home
	home, err := cmd.Flags().GetBool("home")
	if err != nil {
		panic(err)
	}
	if home {
		chdir(os.UserHomeDir())
	}

	// config
	config, err := cmd.Flags().GetBool("config")
	if err != nil {
		panic(err)
	}
	if config {
		chdir(os.UserConfigDir())
	}

	// cache
	cache, err := cmd.Flags().GetBool("cache")
	if err != nil {
		panic(err)
	}
	if cache {
		chdir(os.UserCacheDir())
	}
}

func chdir(dir string, err error) {
	if err != nil {
		panic(err)
	}
	err = os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
