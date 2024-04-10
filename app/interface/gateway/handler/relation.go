package handler

import (
	"context"
	"github.com/Mutezebra/tiktok/app/domain/model/errno"
	"github.com/Mutezebra/tiktok/app/interface/gateway/pack"
	"github.com/Mutezebra/tiktok/app/interface/gateway/rpc"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/kitex_gen/api/relation"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func FollowHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req relation.FollowReq
		if err := c.BindAndValidate(&req); err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(errno.InvalidParamErrno, err))
			return
		}

		req.UID = new(int64)
		*req.UID, _ = strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)

		resp, err := rpc.Follow(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
		return
	}
}

func FollowListHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req relation.GetFollowListReq
		req.UID = new(int64)
		*req.UID, _ = strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)

		resp, err := rpc.GetFollowList(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
		return
	}
}

func FansListHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req relation.GetFansListReq
		req.UID = new(int64)
		*req.UID, _ = strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)

		resp, err := rpc.GetFansList(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
		return
	}
}

func FriendsListHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req relation.GetFriendsListReq
		req.UID = new(int64)
		*req.UID, _ = strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)

		resp, err := rpc.GetFriendsList(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
		return
	}
}
