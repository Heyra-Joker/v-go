/*
@File    :   cmd/cmd.go
@Time    :   2021/12/09 10:00:10
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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Region string

var rootCmd = &cobra.Command{
	Use: "v-go",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("v-go  program, with args: %v\n", args)
	},
}

func init() {

	cobra.OnInitialize(initConfig)

	// add flag
	rootCmd.PersistentFlags().StringP("path", "p", "", "privide local path")
	rootCmd.MarkPersistentFlagRequired("path")
	viper.BindPFlag("path", rootCmd.Flags().Lookup("path"))

	rootCmd.Flags().StringP("show_format", "s", "json", "show video format")
	viper.BindPFlag("show_format", rootCmd.Flags().Lookup("show_format"))
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// initConfig when start run cobra
func initConfig() {
	// do some thing...
}
