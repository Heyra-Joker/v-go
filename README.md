
# Video Compress
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)

A go program that relies on back-end [ffmpeg](http://www.ffmpeg.org/) to process video-related content




## Installation

### v-go
You can download the corresponding v-go binaries via [releases]()

### Mac
Installing [Homebrew](https://brew.sh/)
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```
Installing [ffmepg using brew](https://formulae.brew.sh/formula/ffmpeg)
```bash
brew install ffmpeg
```
Verify ffmepg
```bash
ffmpeg -h
```
### Linux
you can use [install.ffmpeg.sh](script/tools/install.ffmpeg.sh) for simple Installation of ffmepg
```bash
cd <YOUR install.ffmpeg.sh PATH>
chmod +x install.ffmpeg.sh
sudo ./install.ffmpeg.sh
```

### Configure Env
This step can option
```bash
mv <YOUR v-go PATH> /usr/local/bin
```

## Usage
**Show Video Base Info**
```bash
v-go -p <YOUR VIDEO LOCAL PATH>
```
**Compression video**

compress videos according to the [YouTube video specification](https://support.google.com/youtube/answer/2853702?hl=en#zippy=%2Cvariable-bitrate-with-custom-stream-keys-in-live-control-room)
```bash
v-go compress -p <YOUR VIDEO LOCAL PATH>
```

**More information**
```bash
v-go --help
```
## Features

- Delogo clear video logo
- convert video

