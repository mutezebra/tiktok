package video

import (
	"context"
	"fmt"
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

func (s *Service) UploadVideo(ctx context.Context, name string, data []byte) (err error, url string) {
	err, path := s.oss.UploadVideo(ctx, name, data)
	if err != nil {
		return err, path
	}
	url = s.oss.DownloadVideo(ctx, path)
	return nil, url
}

func (s *Service) UploadVideoCover(ctx context.Context, name string, data []byte) (err error, url string) {
	err, path := s.oss.UploadVideoCover(ctx, name, data)
	if err != nil {
		return err, path
	}
	url = s.oss.DownloadVideoCover(ctx, path)
	return nil, url
}

func (s *Service) OssVideoURL(ctx context.Context, vid int64) (string, error) {
	val, err := s.repo.GetValByColumn(ctx, vid, "video_url")
	if err != nil {
		return "", err
	}

	return val, nil
}

func (s *Service) OssCoverURL(ctx context.Context, vid int64) (string, error) {
	val, err := s.repo.GetValByColumn(ctx, vid, "cover_url")
	if err != nil {
		return "", err
	}

	return val, nil
}

func (s *Service) OssVideoName(ctx context.Context, vid int64) (string, error) {
	val, err := s.repo.GetValByColumn(ctx, vid, "video_ext")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d%s", vid, val), nil
}

func (s *Service) OssCoverName(ctx context.Context, vid int64) (string, error) {
	val, err := s.repo.GetValByColumn(ctx, vid, "cover_ext")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d%s", vid, val), nil
}

func (s *Service) VideoFeed(name string) ([]byte, error) {
	return s.oss.VideoFeed(name)
}
