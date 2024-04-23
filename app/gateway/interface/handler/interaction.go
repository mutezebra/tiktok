package handler

import (
	"context"
	"strconv"

	"github.com/Mutezebra/tiktok/app/gateway/domain/model"
	"github.com/Mutezebra/tiktok/app/gateway/interface/pack"
	"github.com/Mutezebra/tiktok/app/gateway/interface/rpc"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/kitex_gen/api/interaction"
)

func LikeHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req interaction.LikeReq
		if err := c.BindAndValidate(&req); err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(model.InvalidParamErrno, err))
			return
		}

		req.UID = new(int64)
		*req.UID, _ = strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)

		resp, err := rpc.Like(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
	}
}

func DisLikeHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req interaction.DislikeReq
		if err := c.BindAndValidate(&req); err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(model.InvalidParamErrno, err))
			return
		}

		req.UID = new(int64)
		*req.UID, _ = strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)

		resp, err := rpc.DisLike(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
	}
}

func LikeListHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req interaction.LikeListReq
		if err := c.BindAndValidate(&req); err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(model.InvalidParamErrno, err))
			return
		}

		req.UID = new(int64)
		*req.UID, _ = strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)

		resp, err := rpc.LikeList(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
	}
}

func CommentHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req interaction.CommentReq
		if err := c.BindAndValidate(&req); err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(model.InvalidParamErrno, err))
			return
		}

		req.UID = new(int64)
		*req.UID, _ = strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)

		resp, err := rpc.Comment(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
	}
}

func CommentListHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req interaction.CommentListReq
		if err := c.BindAndValidate(&req); err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(model.InvalidParamErrno, err))
			return
		}

		resp, err := rpc.CommentList(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
	}
}

func DeleteCommentHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req interaction.DeleteCommentReq
		if err := c.BindAndValidate(&req); err != nil {
			pack.SendFailedResponse(c, pack.ReturnError(model.InvalidParamErrno, err))
			return
		}

		req.UID = new(int64)
		*req.UID, _ = strconv.ParseInt(string(c.GetHeader(consts.HeaderUserIdKey)), 10, 64)

		resp, err := rpc.DeleteComment(ctx, &req)
		if err != nil {
			pack.SendFailedResponse(c, err)
			return
		}

		pack.SendResponse(c, resp)
	}
}
