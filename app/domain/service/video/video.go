package video

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/Mutezebra/tiktok/app/domain/model"
	"github.com/Mutezebra/tiktok/app/domain/repository"
	"github.com/Mutezebra/tiktok/pkg/snowflake"
)

type Service struct {
	repo  repository.VideoRepository
	cache repository.VideoCacheRepository
	oss   model.OSS
}

func NewService(repo repository.VideoRepository, cache repository.VideoCacheRepository, oss model.OSS) *Service {
	return &Service{
		repo:  repo,
		cache: cache,
		oss:   oss,
	}
}

func (s *Service) IsVideo(filename string) bool {
	ext := filepath.Ext(filename)
	imageExts := []string{".mp4", ".avi", ".mov"}
	for _, imageExt := range imageExts {
		if strings.EqualFold(ext, imageExt) {
			return true
		}
	}
	return false
}

func (s *Service) IsImage(filename string) bool {
	ext := filepath.Ext(filename)
	imageExts := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".tiff"}
	for _, imageExt := range imageExts {
		if strings.EqualFold(ext, imageExt) {
			return true
		}
	}
	return false
}

func (s *Service) GenerateID() int64 {
	return snowflake.GenerateID()
}

func (s *Service) UploadVideo(ctx context.Context, name string, data []byte) (err error, path string) {
	return s.oss.UploadVideo(ctx, name, data)
}

func (s *Service) DownloadVideo(ctx context.Context, path string) string {
	return s.oss.DownloadVideo(ctx, path)
}

func (s *Service) VideoFeed(path string) ([]byte, error) {
	return s.oss.VideoFeed(path)
}

func (s *Service) UploadVideoCover(ctx context.Context, name string, data []byte) (err error, path string) {
	return s.oss.UploadVideoCover(ctx, name, data)
}

func (s *Service) DownloadVideoCover(ctx context.Context, path string) string {
	return s.oss.DownloadVideoCover(ctx, path)
}
