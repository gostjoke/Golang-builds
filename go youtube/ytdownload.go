package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func main() {
	client := youtube.Client{}
	var videoURL string
	fmt.Print("請輸入影片網址：")
	fmt.Scanln(&videoURL)

	video, err := client.GetVideo(videoURL)
	if err != nil {
		panic(err)
	}

	// 找出最適合的格式（這邊挑選 MP4 且包含音訊的）
	var format *youtube.Format
	for _, f := range video.Formats {
		if f.MimeType == "video/mp4" && f.AudioChannels > 0 {
			format = &f
			break
		}
	}

	if format == nil {
		panic("找不到適合的格式")
	}

	stream, _, err := client.GetStream(video, format)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}

	fmt.Println("影片下載完成")
}
