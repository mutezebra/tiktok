package usecase

import (
	"context"
	"github.com/Mutezebra/tiktok/app/usecase/pack"
	"github.com/Mutezebra/tiktok/pkg/log"
	"github.com/pkg/errors"
	"time"

	"github.com/Mutezebra/tiktok/app/domain/repository"
	"github.com/Mutezebra/tiktok/app/domain/service/user"
	idl "github.com/Mutezebra/tiktok/kitex_gen/api/user"
)

type UserCase struct {
	repo    repository.UserRepository
	service *user.Service
}

func NewUseUseCase(repo repository.UserRepository, srv *user.Service) *UserCase {
	return &UserCase{
		repo:    repo,
		service: srv,
	}
}

type userDTO struct {
	id             int64
	username       string
	passwordDigest string
	email          string
}

func (u *UserCase) Register(ctx context.Context, req *idl.RegisterReq) (r *idl.RegisterResp, err error) {
	dto := userDTO{}
	dto.username = *req.UserName
	dto.email = *req.Email

	dto.id, err = u.service.GenerateID()
	if err != nil {
		log.LogrusObj.Error(errors.WithMessage(err, "failed generate id"))
		return
	}

	dto.passwordDigest, err = u.service.EncryptPassword(*req.Password)
	if err != nil {
		log.LogrusObj.Error(errors.WithMessage(err, "encrypt password failed"))
		return
	}
	
	cCtx, cn := context.WithTimeout(ctx, 1*time.Second)
	defer cn()
	err = u.repo.CreateUser(cCtx, toUser(&dto))
	if err != nil {
		log.LogrusObj.Error(errors.WithMessage(err, "encrypt password failed"))
		return
	}

	r = new(idl.RegisterResp)
	r.Base = pack.Success
	return r, nil
}

func (u *UserCase) Login(ctx context.Context, req *idl.LoginReq) (r *idl.LoginResp, err error) {
	return nil, err
}
func (u *UserCase) Info(ctx context.Context, req *idl.InfoReq) (r *idl.InfoResp, err error) {
	return nil, err
}
func (u *UserCase) UploadAvatar(ctx context.Context, req *idl.UploadAvatarReq) (r *idl.UploadAvatarResp, err error) {
	return nil, err
}
func (u *UserCase) TotpQrcode(ctx context.Context, req *idl.TotpQrcodeReq) (r *idl.TotpQrcodeResp, err error) {
	return nil, err
}
func (u *UserCase) EnableTotp(ctx context.Context, req *idl.EnableTotpReq) (r *idl.EnableTotpResp, err error) {
	return nil, err
}

func toUser(dto *userDTO) *repository.User {
	return &repository.User{
		ID:             dto.id,
		UserName:       dto.username,
		Email:          dto.email,
		PasswordDigest: dto.passwordDigest,
		Gender:         -1,
		Avatar:         "",
		Fans:           0,
		Follows:        0,
		TotpEnable:     false,
		TotpSecret:     "",
		CreateAt:       time.Now().Unix(),
		UpdateAt:       time.Now().Unix(),
		DeleteAt:       0,
	}
}
