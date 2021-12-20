/*
@File    :   compress/compress.go
@Time    :   2021/12/08 14:11:23
@Author  :   Joker
@Contact :   jokermonster030@gmail.com
@Desc    :   None
🤡
code is far away from bugs with the god animal protecting
I love animals. They taste delicious.
🤡
*/

package compress

import (
	"github.com/Heyra-Joker/v-go/pkg/videoInfo"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

type Video struct {
	path        string
	resolutions string
	frameRate   uint
	bitRate     uint
	videoStream gjson.Result
	output      string
	coverage    bool
}

// New video compress struct to set some variable
func (v *Video) New() {
	v.path = viper.GetString("path")
	v.resolutions = viper.GetString("resolutions")
	v.frameRate = viper.GetUint("frame_rate")
	v.bitRate = viper.GetUint("bit_rate")
	v.coverage = viper.GetBool("coverage")

	// set video info
	vInfoParse := gjson.Parse(videoInfo.ShowVideoInfo(v.path, "json"))
	v.videoStream = vInfoParse.Get("streams").Array()[0]

	// set output
	v.output = viper.GetString("output")
}

func (v *Video) Compress() {
	// verification some condition before compressing
	v.PreCompress()

	// start compressing
	v.Compressing()
}
