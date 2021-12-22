/*
@author:Joker
@file: preCompress.go
@time: 2021/12/13/11:27 上午
@blog: https://github.com/joker-heyra
@description: pre-compression before compress

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/

package compress

import (
	"fmt"
	"os"

	"github.com/Heyra-Joker/v-go/pkg/util"
	"github.com/Heyra-Joker/v-go/pkg/youTubeNormal"
	"github.com/fatih/color"
)

// PreCompress Verify before compressing
func (v Video) PreCompress() {

	// get YouTube specification
	key := fmt.Sprintf("%s@%d", v.resolutions, v.frameRate)
	targetNormal, ok := youTubeNormal.GetYouTubeNormal(key)
	if !ok {
		msg := fmt.Sprintf("YouTube video specification not suggestion resolution: %s, fps: %d", v.resolutions, v.frameRate)
		util.CobraErr(msg)
	}

	// coded_width < coded_height (Vertical video)
	var codeWidth, codeHeight = v.codeWidth, v.codeHeight
	var targetWidth, targetHeight = targetNormal.GetWidth(), targetNormal.GetHeight()
	if !v.isHorizontal {
		codeWidth = v.codeHeight
		codeHeight = v.codeWidth
	}

	// non-acceptance:
	// any situation video Direction is Horizontal
	if targetWidth > codeWidth || targetHeight > codeHeight {
		msg := `video compression can only be done from high resolution to low resolution, your target (%dx%d) higher current (%dx%d)`
		msg = fmt.Sprintf(msg, targetNormal.GetWidth(), targetNormal.GetHeight(), codeWidth, codeHeight)
		util.CobraErr(msg)
	}

	if targetWidth == codeWidth && targetHeight == codeHeight {
		msg := `your target (%dx%d) resolution equal current (%dx%d), no action required`
		color.Yellow(msg, targetNormal.GetWidth(), targetNormal.GetHeight(), codeWidth, codeHeight)
		os.Exit(0)
	}

	if v.bitRate < targetNormal.GetBitRateMin() || v.bitRate > targetNormal.GetBitRateMax() {
		msg := `YouTube video specification suggest bit rate range %d~%d when quality equal %s, current bit rate: %d`
		msg = fmt.Sprintf(msg, targetNormal.GetBitRateMin(), targetNormal.GetBitRateMax(), key, v.bitRate)
		util.CobraErr(msg)
	}
	// non-acceptance

}
