package handler

import (
	"context"
	"github.com/Mutezebra/tiktok/app/usecase"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/pkg/log"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"strconv"
)

func ChatHandler() app.HandlerFunc {
	var upgrader = websocket.HertzUpgrader{} // use default options

	return func(ctx context.Context, c *app.RequestContext) {
		UID, _ := strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)

		toS := c.Query("to")
		to, err := strconv.ParseInt(toS, 10, 64)
		if err != nil {
			c.JSON(200, "to format error")
			return
		}
		err = upgrader.Upgrade(c, usecase.ChatHandler(ctx, UID, to))
		if err != nil {
			log.LogrusObj.Errorf("in front failed %v", err)
			c.JSON(200, err)
			return
		}
	}
}
