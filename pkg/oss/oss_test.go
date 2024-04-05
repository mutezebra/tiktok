package oss

import (
	"context"
	"fmt"
	"testing"

	"github.com/Mutezebra/tiktok/config"
)

func TestModel_DownloadAvatar(t *testing.T) {
	config.InitConfig()
	InitOSS()
	model := NewOssModel()
	url := model.DownloadVideo(context.Background(), "static/videos/covers/1749124012052480.png")
	if url == "" {
		fmt.Println("is empty")
	}
	fmt.Println(url)
}
