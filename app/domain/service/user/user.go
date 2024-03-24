package user

import (
	"golang.org/x/crypto/bcrypt"
	"net/mail"

	"github.com/Mutezebra/tiktok/app/domain/repository"
	"github.com/Mutezebra/tiktok/app/domain/service/user/pkg/snowflake"
)

type Service struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) *Service {
	return &Service{repo: repo}
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
