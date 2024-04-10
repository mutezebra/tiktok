package handler

import (
	"context"
	"github.com/Mutezebra/tiktok/app/domain/model/errno"
	"github.com/Mutezebra/tiktok/app/interface/gateway/pack"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"

	"github.com/Mutezebra/tiktok/app/usecase"
	"github.com/Mutezebra/tiktok/consts"
)

func ChatHandler() app.HandlerFunc {
	var upgrader = websocket.HertzUpgrader{} // use default options

	return func(ctx context.Context, c *app.RequestContext) {
		UID, _ := strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)

		toS := c.Query("to")
		to, err := strconv.ParseInt(toS, 10, 64)
		if err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(errno.InvalidParamErrno, err))
			return
		}

		err = upgrader.Upgrade(c, usecase.ChatHandler(ctx, UID, to))
		if err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(errno.InternalServerErrorErrno, err))
			return
		}
	}
}
