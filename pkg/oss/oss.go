package oss

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Model struct {
}

func NewOssModel() *Model {
	return &Model{}
}

// UploadAvatar from the oss
func (m *Model) UploadAvatar(ctx context.Context, name string, data []byte) (err error, path string) {
	mac, bucket, _ := getOSS()
	path = conf.AvatarPath + name

	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, path),
	}
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Region: &storage.ZoneHuadong,
	}
	formLoader := storage.NewFormUploader(&cfg)

	ret := &PutRet{}
	putExtra := &storage.PutExtra{}

	length := int64(len(data))

	err = formLoader.Put(ctx, ret, upToken, path, bytes.NewReader(data), length, putExtra)
	return err, path
}

func (m *Model) DownloadAvatar(ctx context.Context, path string) string {
	_, _, domain := getOSS()

	publicURL := storage.MakePublicURLv2(domain, path)
	return publicURL
}
