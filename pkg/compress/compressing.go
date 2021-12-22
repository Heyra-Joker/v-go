/*
@author:Joker
@file: compressing.go
@time: 2021/12/13/2:17 下午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/

package compress

import (
	"fmt"
	"github.com/Heyra-Joker/v-go/pkg/commder"
	"github.com/Heyra-Joker/v-go/pkg/util"
	"github.com/Heyra-Joker/v-go/pkg/youTubeNormal"
	"github.com/fatih/color"
)

// Specify the Height To Retain the Aspect Ratio
// https://dtbaker.net/blog/ffmpeg-width-or-height-not-divisible-by-2/
//  tell ffmpeg to choose a number that’s divisible by 2 (instead of 1)
const compressCmdPre = "ffmpeg -i %s -r %d -b:v %dk -vf scale=-2:%d -crf 28 %s"

// Compressing video
func (v Video) Compressing() {
	key := fmt.Sprintf("%s@%d", v.resolutions, v.frameRate)
	targetNormal, _ := youTubeNormal.GetYouTubeNormal(key)

	cmd := fmt.Sprintf(compressCmdPre, v.path, v.frameRate, v.bitRate, targetNormal.GetHeight(), v.output)
	if v.coverage {
		cmd += " -y"
	}

	result, err := commder.Command(cmd)
	if err != nil {
		msg := fmt.Sprintf("sorry, compress video failed,\nerr: %s", err.Error())
		util.CobraErr(msg)
	}

	if !result.Success {
		msg := fmt.Sprintf("sorry, compress video failed,\nerr: %s", result.StdErr)
		util.CobraErr(msg)
	}

	color.Green("compress %s success, you can use v-go -p %s see information", v.path, v.output)
}
