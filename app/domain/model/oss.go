package model

import (
	"context"
)

type OSS interface {
	UploadAvatar(ctx context.Context, name string, data []byte) (err error, path string)
	DownloadAvatar(ctx context.Context, key string) string
}
