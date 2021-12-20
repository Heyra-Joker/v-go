/*
@author:Joker
@file: compressCommand.go
@time: 2021/12/13/10:27 上午
@blog: https://github.com/joker-heyra
@description: judgement compress command

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/

package judgement

import (
	"errors"
	"github.com/Heyra-Joker/v-go/pkg/util"
	"github.com/fatih/color"
	"strings"
)

type Compress struct {
	Resolutions string
	FrameRate   uint
	Output      string
}

var allowResolutions = []string{"240p", "360p", "480p", "720p", "1080p", "1440p", "4k", "2160p"}
var allowFrameRate = []uint{30, 60}

func (c Compress) verify() error {
	exists := util.ItemContainer(c.Resolutions, allowResolutions)
	if !exists {
		return errors.New(color.RedString("resolutions accept: %s", strings.Join(allowResolutions, ", ")))
	}

	exists = util.ItemContainer(c.FrameRate, allowFrameRate)
	if !exists {
		return errors.New(color.RedString("frame_rate illegal"))
	}

	return nil
}
