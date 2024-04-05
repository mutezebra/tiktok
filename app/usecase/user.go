package usecase

import (
	"context"
	"time"

	"github.com/Mutezebra/tiktok/app/domain/model/errno"
	"github.com/Mutezebra/tiktok/app/usecase/pack"
	"github.com/Mutezebra/tiktok/pkg/utils"

	"github.com/Mutezebra/tiktok/app/domain/repository"
	userService "github.com/Mutezebra/tiktok/app/domain/service/user"
	idl "github.com/Mutezebra/tiktok/kitex_gen/api/user"
)

type UserCase struct {
	repo    repository.UserRepository
	service *userService.Service
}

func NewUseUseCase(repo repository.UserRepository, srv *userService.Service) *UserCase {
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
	dto.username = req.GetUserName()
	dto.id = u.service.GenerateID()

	if dto.email, err = u.service.VerifyEmail(req.GetEmail()); err != nil {
		return nil, pack.ReturnError(errno.EmailFormatError, err)
	}

	if dto.passwordDigest, err = u.service.EncryptPassword(req.GetPassword()); err != nil {
		return nil, pack.ReturnError(errno.EncryptPasswordError, err)
	}

	if err = u.repo.CreateUser(ctx, dtoU2Repo(&dto)); err != nil {
		return nil, pack.ReturnError(errno.DatabaseCreateUserError, err)
	}

	r = new(idl.RegisterResp)

	return r, nil
}

func (u *UserCase) Login(ctx context.Context, req *idl.LoginReq) (r *idl.LoginResp, err error) {
	passwordDigest, id, err := u.repo.GetPasswordAndIDByName(ctx, req.GetUserName())
	if err != nil {
		return nil, pack.ReturnError(errno.GetPasswordFromDatabaseError, err)
	}

	ok := u.service.CheckPassword(req.GetPassword(), passwordDigest)
	if !ok {
		return nil, pack.ReturnError(errno.CheckPasswordError, nil)
	}

	aToken, rToken, err := utils.GenerateToken(req.GetUserName(), id)
	if err != nil {
		return nil, pack.ReturnError(errno.GenerateTokenError, err)
	}

	r = new(idl.LoginResp)
	r.SetAccessToken(&aToken)
	r.SetRefreshToken(&rToken)

	return r, nil
}

func (u *UserCase) Info(ctx context.Context, req *idl.InfoReq) (r *idl.InfoResp, err error) {
	user, err := u.repo.UserInfoByID(ctx, *req.UID)
	if err != nil {
		return nil, pack.ReturnError(errno.GetUserInfoError, err)
	}

	r = new(idl.InfoResp)
	r.Data = repoU2IDL(*user)

	return r, nil
}

func (u *UserCase) UploadAvatar(ctx context.Context, req *idl.UploadAvatarReq) (r *idl.UploadAvatarResp, err error) {
	ok, avatar := u.service.AvatarName(*req.FileName, *req.UID)
	if !ok {
		return nil, pack.ReturnError(errno.GetAvatarNameError, nil)
	}

	err, path := u.service.OSS.UploadAvatar(ctx, avatar, req.Avatar)
	if err != nil {
		return nil, pack.ReturnError(errno.OssUploadAvatarError, err)
	}

	url := u.service.OSS.DownloadAvatar(ctx, path)
	if url == "" {

	}

	if err = u.repo.UpdateUserAvatar(ctx, path, *req.UID); err != nil {
		return nil, pack.ReturnError(errno.DatabaseUpdateUserAvatarError, err)
	}

	r = new(idl.UploadAvatarResp)
	return r, nil
}

func (u *UserCase) DownloadAvatar(ctx context.Context, req *idl.DownloadAvatarReq) (r *idl.DownloadAvatarResp, err error) {
	path, err := u.repo.GetUserAvatar(ctx, req.GetUID())
	if err != nil {
		return nil, pack.ReturnError(errno.DatabaseGetUserAvatarError, err)
	}

	url := u.service.DownloadAvatar(ctx, path)
	if url == "" {
		return nil, pack.ReturnError(errno.OssDownloadAvatarError, nil)
	}

	r = new(idl.DownloadAvatarResp)
	r.SetURL(&url)
	return r, nil
}

func (u *UserCase) TotpQrcode(ctx context.Context, req *idl.TotpQrcodeReq) (r *idl.TotpQrcodeResp, err error) {
	secret, png, err := u.service.GenerateTotp(req.GetUserName())
	if err != nil {
		return nil, pack.ReturnError(errno.GenerateTotpError, err)
	}

	err = u.repo.UpdateTotpSecret(ctx, req.GetUID(), secret)
	if err != nil {
		return nil, pack.ReturnError(errno.DatabaseUpdateTotpSecretError, err)
	}

	r = new(idl.TotpQrcodeResp)
	r.SetQrcode(&png)
	return r, nil
}

func (u *UserCase) EnableTotp(ctx context.Context, req *idl.EnableTotpReq) (r *idl.EnableTotpResp, err error) {
	secret, err := u.repo.GetTotpSecret(ctx, req.GetUID())
	if err != nil {
		return nil, pack.ReturnError(errno.DatabaseGetTotpSecretError, err)
	}

	ok := u.service.VerifyOtp(req.GetCode(), secret)
	if !ok {
		return nil, pack.ReturnError(errno.VerifyOtpCodeError, err)
	}

	err = u.repo.UpdateTotpStatus(ctx, true, req.GetUID())
	if err != nil {
		return nil, pack.ReturnError(errno.DatabaseUpdateTotpStatusError, err)
	}

	r = new(idl.EnableTotpResp)
	return r, nil
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
		TotpStatus: &user.TotpEnable,
	}
}
