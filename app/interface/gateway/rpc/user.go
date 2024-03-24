package rpc

import (
	"context"
	"github.com/Mutezebra/tiktok/kitex_gen/api/user"
)

func Register(ctx context.Context, req *user.RegisterReq) (r *user.RegisterResp, err error) {
	r, err = UserClient.Register(ctx, req)
	return r, err
}

func Login(ctx context.Context, req *user.LoginReq) (r *user.LoginResp, err error) {
	r, err = UserClient.Login(ctx, req)
	return r, err
}

func Info(ctx context.Context, req *user.InfoReq) (r *user.InfoResp, err error) {
	r, err = UserClient.Info(ctx, req)
	return r, err
}

func UploadAvatar(ctx context.Context, req *user.UploadAvatarReq) (r *user.UploadAvatarResp, err error) {
	r, err = UserClient.UploadAvatar(ctx, req)
	return r, err
}

func TotpQrcode(ctx context.Context, req *user.TotpQrcodeReq) (r *user.TotpQrcodeResp, err error) {
	r, err = UserClient.TotpQrcode(ctx, req)
	return r, err
}

func EnableTotp(ctx context.Context, req *user.EnableTotpReq) (r *user.EnableTotpResp, err error) {
	r, err = UserClient.EnableTotp(ctx, req)
	return r, err
}
