package user

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/Mutezebra/tiktok/app/domain/repository"
	"github.com/Mutezebra/tiktok/app/domain/service/user/pkg/snowflake"
)

type Service struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) *Service {
	return &Service{repo: repo}
}

func (srv *Service) GenerateID() (int64, error) {
	id := snowflake.GenerateID()
	if id == 0 {
		err := errors.New("generate id failed")
		return 0, errors.Wrap(err, "")
	}
	return id, nil
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
