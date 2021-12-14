/*
@author:Joker
@file: youTubeNormal.go
@time: 2021/12/13/11:34 上午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/

package youTubeNormal

import "sync"

var once sync.Once

type YouTubeNormal interface {
	GetWidth() uint
	GetHeight() uint
	GetFps() uint8
	GetBitRateMin() uint
	GetBitRateMax() uint
}

type VideoNormal struct {
	width      uint
	height     uint
	fps        uint8
	bitRateMin uint
	bitRateMax uint
}

func (n VideoNormal) GetWidth() uint {
	return n.width
}

func (n VideoNormal) GetHeight() uint {
	return n.height
}

func (n VideoNormal) GetFps() uint8 {
	return n.fps
}

func (n VideoNormal) GetBitRateMin() uint {
	return n.bitRateMin
}

func (n VideoNormal) GetBitRateMax() uint {
	return n.bitRateMax
}

var youTubeVideoNormals map[string]YouTubeNormal

// Register Youtube video normal map
func Register(resolution string, width, height, bitRateMin, bitRateMax uint, fps uint8) {
	once.Do(func() {
		youTubeVideoNormals = make(map[string]YouTubeNormal, 0)
	})

	youTubeVideoNormals[resolution] = VideoNormal{
		width:      width,
		height:     height,
		fps:        fps,
		bitRateMin: bitRateMin,
		bitRateMax: bitRateMax,
	}

}

// RegisterDefaultYouTubeNormals
// register default
func RegisterDefaultYouTubeNormals() {
	Register("240p@30", 426, 240, 300, 700, 30)
	Register("360p@30", 640, 360, 400, 1000, 30)
	Register("480p@30", 854, 480, 500, 2000, 30)

	Register("720p@30", 1280, 720, 1500, 4000, 30)
	Register("720p@60", 1280, 720, 2250, 6000, 60)

	Register("1080p@30", 1920, 1080, 3000, 6000, 30)
	Register("1080p@60", 1920, 1080, 4500, 9000, 60)

	Register("1440p@30", 2560, 1440, 6000, 13000, 30)
	Register("1440p@60", 2560, 1440, 9000, 18000, 60)

	Register("4k@30", 3840, 2160, 13000, 34000, 30)
	Register("4k@60", 3840, 2160, 20000, 51000, 60)

	Register("2160p@30", 3840, 2160, 13000, 34000, 30)
	Register("2160p@60", 3840, 2160, 20000, 51000, 60)

}

// GetYouTubeNormal get specification from YouTube normal map
// key = "resolution@fps"
func GetYouTubeNormal(key string) (YouTubeNormal, bool) {
	normal, ok := youTubeVideoNormals[key]
	return normal, ok
}
