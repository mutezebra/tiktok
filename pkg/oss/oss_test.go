package oss

import (
	"context"
	"fmt"
	"github.com/Mutezebra/tiktok/config"
	"testing"
)

func TestModel_DownloadAvatar(t *testing.T) {
	config.InitConfig()
	InitOSS()
	model := NewOssModel()
	_, url := model.DownloadAvatar(context.Background(), "static/avatars/968796674723840.png")
	if url == "" {
		fmt.Println("is empty")
	}
	fmt.Println(url)
}
