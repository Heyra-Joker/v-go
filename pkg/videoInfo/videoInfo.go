/*
@File    :   videoInfo/videoInfo.go
@Time    :   2021/12/10 10:52:15
@Author  :   Joker
@Contact :   jokermonster030@gmail.com
@Desc    :   None
🤡
code is far away from bugs with the god animal protecting
I love animals. They taste delicious.
🤡
*/

package videoInfo

import (
	"errors"
	"fmt"
	"github.com/Heyra-Joker/video-compress/pkg/util"
	"io/ioutil"

	"github.com/Heyra-Joker/video-compress/pkg/commder"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// -select_streams v: video a:audio
const vInfoCmdPre = "ffprobe -print_format %s -show_streams -select_streams v -show_format %s"
const failedInfoPre = "sorry, path: '%s' get video info failed.\nmore info:\n%s"

// ShowVideoInfo show video use backend 'ffmpeg'
func ShowVideoInfo(path, showFormat string) string {

	//  show video info if success
	result, err := commder.Command(fmt.Sprintf(vInfoCmdPre, showFormat, path))
	if err != nil {
		cobra.CheckErr(errors.New(color.RedString(failedInfoPre, path, err.Error())))
	}

	if !result.Success {
		cobra.CheckErr(errors.New(color.RedString(failedInfoPre, path, result.StdErr)))
	}

	return result.Stdout
}

// SaveVideoInfo save video info to target file
func SaveVideoInfo(info, savePath string) {
	if savePath == "" {
		fmt.Println(info)
		return
	}

	err := ioutil.WriteFile(savePath, []byte(info), 0666)
	if err != nil {
		msg := "sorry save file error, more information:\n"
		msg = fmt.Sprintf(msg, err.Error())
		util.CobraErr(msg)
	}
}
