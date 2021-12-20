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
	"github.com/Heyra-Joker/v-go/pkg/compress"
	"github.com/Heyra-Joker/v-go/pkg/judgement"
	"github.com/Heyra-Joker/v-go/pkg/youTubeNormal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "compress video",
	Long:  "compression according to the youtube specification:\n https://support.google.com/youtube/answer/2853702?hl=en#zippy=%2Cvariable-bitrate-with-custom-stream-keys-in-live-control-room",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		compress := judgement.Compress{
			Resolutions: viper.GetString("resolutions"),
			FrameRate:   viper.GetUint("frame_rate"),
		}

		if err := judgement.New(compress); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		video := compress.Video{}
		video.New()
		video.Compress()
	},
}

func init() {

	// add flag
	compressCmd.Flags().StringP("resolutions", "r", "720p", "see youtube specification")
	_ = viper.BindPFlag("resolutions", compressCmd.Flags().Lookup("resolutions"))

	compressCmd.Flags().Uint8P("frame_rate", "f", 30, "video frame rate (30, 60) fps")
	_ = viper.BindPFlag("frame_rate", compressCmd.Flags().Lookup("frame_rate"))

	compressCmd.Flags().UintP("bit_rate", "b", 1500, "video bit rate unit(k)")
	_ = viper.BindPFlag("bit_rate", compressCmd.Flags().Lookup("bit_rate"))
	rootCmd.AddCommand(compressCmd)

	// register YouTube video specification
	youTubeNormal.RegisterDefaultYouTubeNormals()
}
