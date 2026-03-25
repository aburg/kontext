/*
Copyright © 2026 Andreas Burg <andreas.burg@gmx.net>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/aburg/kontext/util"
	"github.com/spf13/cobra"
)

// unsetCmd represents the unset command
var unsetCmd = &cobra.Command{
	Use:   "unset <kind>",
	Short: "unset the context of the given kind (and return the latest value)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		util.ChangeWorkdir(cmd)

		kind := args[0]
		path, err := util.FindContext(kind)
		if err != nil {
			return err
		}

		if err = util.PrintFileContent(path); err != nil {
			return err
		}

		if err = util.DeleteFile(path); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(unsetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unsetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unsetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
