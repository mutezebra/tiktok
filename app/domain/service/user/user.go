package user

import (
	"context"
	"fmt"
	"github.com/Mutezebra/tiktok/app/domain/model"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"path/filepath"
	"strings"

	"github.com/Mutezebra/tiktok/app/domain/repository"
	"github.com/Mutezebra/tiktok/pkg/snowflake"
)

type Service struct {
	repo repository.UserRepository
	OSS  model.OSS
}

func NewService(repo repository.UserRepository, oss model.OSS) *Service {
	return &Service{repo: repo, OSS: oss}
}

func (srv *Service) GenerateID() int64 {
	return snowflake.GenerateID()
}

func (srv *Service) EncryptPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "nil", err
	}
	result := string(hashPassword)

	return result, nil
}

func (srv *Service) CheckPassword(password string, passwordDigest string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordDigest), []byte(password)) == nil
}

func (srv *Service) VerifyEmail(email string) (string, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", err
	}
	return email, nil
}

func (srv *Service) UploadAvatar(ctx context.Context, name string, data []byte) (err error, path string) {
	return srv.OSS.UploadAvatar(ctx, name, data)
}

func (srv *Service) DownloadAvatar(ctx context.Context, name string) (url string) {
	return srv.OSS.DownloadAvatar(ctx, name)
}

// AvatarName get the avatar filename
func (srv *Service) AvatarName(filename string, id int64) (ok bool, avatarName string) {
	ext := filepath.Ext(filename)
	imageExts := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".tiff"}
	for _, imageExt := range imageExts {
		if strings.EqualFold(ext, imageExt) {
			ok = true
		}
	}
	if !ok {
		return false, ""
	}
	avatarName = fmt.Sprintf("%d%s", id, ext)
	return true, avatarName
}
