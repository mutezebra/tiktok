package usecase

import (
	"context"
	"github.com/Mutezebra/tiktok/app/domain/model/errno"
	"github.com/Mutezebra/tiktok/app/usecase/pack"
	"github.com/Mutezebra/tiktok/pkg/utils"
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
	dto.id = u.service.GenerateID()

	if dto.email, err = u.service.VerifyEmail(*req.Email); err != nil {
		return nil, pack.ReturnError(errno.EmailFormatError, err)
	}

	if dto.passwordDigest, err = u.service.EncryptPassword(*req.Password); err != nil {
		return nil, pack.ReturnError(errno.EncryptPasswordError, err)
	}

	if err = u.repo.CreateUser(ctx, dtoU2Repo(&dto)); err != nil {
		return nil, pack.ReturnError(errno.DatabaseCreateUserError, err)
	}

	r = new(idl.RegisterResp)
	return r, nil
}

func (u *UserCase) Login(ctx context.Context, req *idl.LoginReq) (r *idl.LoginResp, err error) {
	passwordDigest, id, err := u.repo.GetPasswordAndIDByName(ctx, *req.UserName)
	if err != nil {
		return nil, pack.ReturnError(errno.GetPasswordFromDatabaseError, err)
	}

	ok := u.service.CheckPassword(*req.Password, passwordDigest)
	if !ok {
		return nil, pack.ReturnError(errno.CheckPasswordError, nil)
	}

	aToken, rToken, err := utils.GenerateToken(*req.UserName, id)
	if err != nil {
		return nil, pack.ReturnError(errno.GenerateTokenError, err)
	}

	r = new(idl.LoginResp)
	r.AccessToken = &aToken
	r.RefreshToken = &rToken
	return r, nil
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

// toUser is userDTO to repository.user
func dtoU2Repo(dto *userDTO) *repository.User {
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

func repoU2IDL(user repository.User) *idl.UserInfo {
	return &idl.UserInfo{
		ID:         &user.ID,
		UserName:   &user.UserName,
		Email:      &user.Email,
		Gender:     &user.Gender,
		Avatar:     &user.Avatar,
		Fans:       &user.Fans,
		Follows:    &user.Follows,
		CreateAt:   &user.CreateAt,
		UpdateAt:   &user.UpdateAt,
		TotpStatus: nil,
	}
}
