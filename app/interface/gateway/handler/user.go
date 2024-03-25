package handler

import (
	"context"
	"github.com/Mutezebra/tiktok/consts"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Mutezebra/tiktok/app/domain/model/errno"
	"github.com/Mutezebra/tiktok/app/interface/gateway/pack"
	"github.com/Mutezebra/tiktok/app/interface/gateway/rpc"
	"github.com/Mutezebra/tiktok/kitex_gen/api/user"
)

func UserRegisterHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req user.RegisterReq
		if err := c.BindAndValidate(&req); err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(errno.InvalidParamErrno, err))
			return
		}

		resp, err := rpc.Register(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
		return
	}
}

func UserLoginHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req user.LoginReq
		if err := c.BindAndValidate(&req); err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(errno.InvalidParamErrno, err))
			return
		}

		resp, err := rpc.Login(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
		return
	}
}

func UserInfoHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req user.InfoReq
		if err := c.BindAndValidate(&req); err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(errno.InvalidParamErrno, err))
			return
		}
		if req.UID == nil {
			req.UID = new(int64)
			*req.UID, _ = strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)
		}
		resp, err := rpc.Info(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}
		pack.SendResponse(c, resp)
		return
	}
}

func UploadAvatarHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		fileHeader, err := c.FormFile(consts.FormUserAvatarKey)
		if err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(errno.InternalServerErrorErrno, err))
			return
		}

		if fileHeader.Size > consts.MB*5 {
			pack.SendFailedResponse(c, pack.ReturnError(errno.OutOfLimitAvatarSizeErrno, nil))
			return
		}

		avatar, err := fileHeader.Open()
		if err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(errno.InternalServerErrorErrno, err))
			return
		}

		var req user.UploadAvatarReq
		req.Avatar = make([]byte, fileHeader.Size)
		_, err = avatar.Read(req.Avatar)
		if err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(errno.InternalServerErrorErrno, err))
			return
		}

		req.UID = new(int64)
		*req.UID, _ = strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)
		req.SetFileName(&fileHeader.Filename)

		resp, err := rpc.UploadAvatar(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
		return
	}
}

func DownloadAvatarHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req user.DownloadAvatarReq

		req.UID = new(int64)
		*req.UID, _ = strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)

		resp, err := rpc.DownloadAvatar(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
		return
	}
}

func TotpQRCodeHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req user.TotpQrcodeReq

		req.UID = new(int64)
		*req.UID, _ = strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)

		resp, err := rpc.TotpQrcode(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
		return
	}
}

func EnableTotpHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req user.EnableTotpReq
		if err := c.BindAndValidate(&req); err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(errno.InvalidParamErrno, err))
			return
		}

		req.UID = new(int64)
		*req.UID, _ = strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)

		resp, err := rpc.EnableTotp(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
		return
	}
}
