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
	"github.com/Heyra-Joker/v-go/pkg/judgement"
	"github.com/Heyra-Joker/v-go/pkg/videoInfo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "v-go",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		// judgement root flag
		root := judgement.Root{
			Path:       viper.GetString("path"),
			ShowFormat: viper.GetString("show_format"),
			Output:     viper.GetString("output"),
			Coverage:   viper.GetBool("coverage"),
			Name:       cmd.Name(),
		}
		if err := judgement.New(root); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		info := videoInfo.ShowVideoInfo(viper.GetString("path"), viper.GetString("show_format"))
		videoInfo.SaveVideoInfo(info, viper.GetString("output"))
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	// add flag
	rootCmd.PersistentFlags().StringP("path", "p", "", "provide local path")
	_ = viper.BindPFlag("path", rootCmd.PersistentFlags().Lookup("path"))

	rootCmd.PersistentFlags().StringP("output", "o", "", "output path, add prefix 'vgo' in video name, if not provide\nNote: without 'show video info'")
	_ = viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))

	rootCmd.PersistentFlags().BoolP("coverage", "c", false, "coverage output")
	_ = viper.BindPFlag("coverage", rootCmd.PersistentFlags().Lookup("coverage"))

	rootCmd.Flags().StringP("show_format", "s", "json", "available formats are: csv, flat, ini, json, xml")
	_ = viper.BindPFlag("show_format", rootCmd.Flags().Lookup("show_format"))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// initConfig when start run cobra
func initConfig() {
	// do some thing...
}
