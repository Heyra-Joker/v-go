/*
@File    :   cmd/version.go
@Time    :   2021/12/09 15:19:55
@Author  :   Joker
@Contact :   jokermonster030@gmail.com
@Desc    :   None
🤡
code is far away from bugs with the god animal protecting
I love animals. They taste delicious.
🤡
*/

package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var version = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show current version",
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("v-go version %s", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
