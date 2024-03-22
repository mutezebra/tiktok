package handler

import (
	"context"
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
			pack.SendFailedResponse(c, errno.InvalidParamErrno)
			return
		}

		resp, err := rpc.Register(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, errno.InternalServerErrorErrno)
			return
		}

		pack.SendResponse(c, resp)
		return
	}
}
