package middleware

import (
    "context"
    "strconv"

    "github.com/cloudwego/hertz/pkg/app"

    "github.com/mutezebra/tiktok/gateway/domain/model"
    "github.com/mutezebra/tiktok/gateway/interface/pack"
    "github.com/mutezebra/tiktok/pkg/consts"
    "github.com/mutezebra/tiktok/pkg/errno"
    "github.com/mutezebra/tiktok/pkg/utils"
)

func JWT() app.HandlerFunc {
    return func(ctx context.Context, c *app.RequestContext) {
        aToken := string(c.GetHeader(consts.HeaderAccessTokenKey))
        rToken := string(c.GetHeader(consts.HeaderRefreshTokenKey))
        if aToken == "" || rToken == "" {
            pack.SendFailedResponse(c, model.UnauthorizedErrno)
            c.Abort()
            return
        }

        claim, err, count := utils.CheckAndUpdateToken(aToken, rToken)
        if err != nil {
            pack.SendFailedResponse(c, errno.WithError(model.UnauthorizedErrno, err))
            c.Abort()
            return
        }
        c.Request.SetHeader(consts.HeaderUserIdKey, strconv.FormatInt(claim.ID, 10))
        c.Request.SetHeader(consts.HeaderUserNameKey, claim.UserName)

        c.Header(consts.HeaderTokenUpdateCountKey, strconv.Itoa(count))
        if count == 0 {
            c.Next(ctx)
            return
        }
        c.Header(consts.HeaderAccessTokenKey, claim.AccessToken)
        c.Next(ctx)
    }
}
