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

type Compresser interface {
	Run()
}

type VideoPrams struct {
	Path      string
	BitRate   string
	FrameRate uint8
	Width     uint16
	Height    uint16
}

// Compression video
func (p VideoPrams) Compression() {

}
