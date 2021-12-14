/*
@File    :   cmd/delogo.go
@Time    :   2021/12/09 16:40:09
@Author  :   Joker
@Contact :   jokermonster030@gmail.com
@Desc    :   None
🤡
code is far away from bugs with the god animal protecting
I love animals. They taste delicious.
🤡
*/

package cmd

import "github.com/spf13/cobra"

var deLogoCmd = &cobra.Command{
	Use:   "delogo",
	Short: "remove video watermark",
	Long:  "using ffmpeg delogo function to remove video watermark, note:\nCurrently only one orientation is supported",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	deLogoCmd.Flags().StringP("position", "P", "right", "[lt|rt|lb|rb]")

	rootCmd.AddCommand(deLogoCmd)
}
