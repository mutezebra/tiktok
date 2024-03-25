package middleware

import (
	"context"
	"github.com/Mutezebra/tiktok/app/domain/model/errno"
	"github.com/Mutezebra/tiktok/app/interface/gateway/pack"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func JWT() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		aToken := string(c.GetHeader(consts.HeaderAccessTokenKey))
		rToken := string(c.GetHeader(consts.HeaderRefreshTokenKey))
		if aToken == "" || rToken == "" {
			pack.SendFailedResponse(c, errno.UnauthorizedErrno)
			c.Abort()
			return
		}

		claim, err, count := utils.CheckAndUpdateToken(aToken, rToken)
		if err != nil {
			pack.SendFailedResponse(c, errno.UnauthorizedErrno)
			c.Abort()
			return
		}
		c.Request.SetHeader(consts.HeaderUserIdKey, strconv.FormatInt(claim.ID, 10))
		c.Request.SetHeader(consts.HeaderUserNameKey, claim.UserName)

		c.Header(consts.HeaderTokenUpdateCountKey, strconv.Itoa(count))
		if count == 0 {
			c.Next(ctx)
		}
		c.Header(consts.HeaderAccessTokenKey, claim.AccessToken)
		c.Next(ctx)
	}
}
