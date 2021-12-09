/*
@File    :   cmd/compress.go
@Time    :   2021/12/09 15:44:44
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
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "compress video",
	Long:  "compression according to the youtube specification:\n https://support.google.com/youtube/answer/2853702?hl=en#zippy=%2Cvariable-bitrate-with-custom-stream-keys-in-live-control-room",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("running...", args)
	},
}

func init() {

	// add flag
	compressCmd.Flags().StringP("resolutions", "r", "720p", "see youtube specification")
	viper.BindPFlag("resolutions", compressCmd.Flags().Lookup("resolutions"))

	compressCmd.Flags().Int8P("frame_rate", "f", 30, "video frame rate")
	viper.BindPFlag("frame_rate", compressCmd.Flags().Lookup("frame_rate"))

	rootCmd.AddCommand(compressCmd)
}
