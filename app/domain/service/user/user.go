package user

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/mail"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/Mutezebra/tiktok/app/domain/model"

	"github.com/Mutezebra/tiktok/app/domain/repository"
	"github.com/Mutezebra/tiktok/pkg/snowflake"
)

type Service struct {
	Repo repository.UserRepository
	OSS  model.OSS
	MFA  model.MFA
}

func NewService(service *Service) *Service {
	if service.Repo == nil {
		panic("user service.Repo should not be nil")
	}
	if service.OSS == nil {
		panic("user service.OSS should not be nil")
	}
	if service.MFA == nil {
		panic("user service.MFA should not be nil")
	}
	return service
}

func (srv *Service) GenerateID() int64 {
	return snowflake.GenerateID()
}

func (srv *Service) EncryptPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "nil", errors.Wrap(err, "failed to encrypt password")
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
		return "", errors.Wrap(err, "invalid email format")
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

func (srv *Service) GenerateTotp(userName string) (secret string, base64 string, err error) {
	return srv.MFA.GenerateTotp(userName)
}

func (srv *Service) VerifyOtp(token string, secret string) bool {
	return srv.MFA.VerifyOtp(token, secret)
}
