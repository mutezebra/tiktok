package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/Mutezebra/tiktok/app/domain/model/errno"
	"github.com/Mutezebra/tiktok/app/domain/repository"
	interactionService "github.com/Mutezebra/tiktok/app/domain/service/interaction"
	"github.com/Mutezebra/tiktok/app/interface/gateway/pack"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/kitex_gen/api/interaction"
	"github.com/Mutezebra/tiktok/kitex_gen/api/video"
)

type InteractionCase struct {
	repo    repository.InteractionRepository
	service *interactionService.Service
}

func NewInteractionCase(repo repository.InteractionRepository, service *interactionService.Service) *InteractionCase {
	return &InteractionCase{
		repo:    repo,
		service: service,
	}
}

type commentDTO struct {
	cid     int64
	uid     int64
	vid     int64
	rootID  int64
	replyID int64
	content string
}

func (i *InteractionCase) Like(ctx context.Context, req *interaction.LikeReq) (r *interaction.LikeResp, err error) {
	var exist bool
	switch req.GetActionType() {
	case consts.LikeCommentActionKey:
		exist, err = i.repo.WhetherCommentLikeItemExist(ctx, req.GetUID(), req.GetCommentID())
		if err != nil {
			return nil, pack.ReturnError(errno.DatabaseWhetherCommentLikeItemExistError, err)
		}
		if exist {
			return nil, pack.ReturnError(errno.CommentAlreadyLikedError, nil)
		}

		exist, err = i.repo.WhetherCommentExist(ctx, req.GetCommentID())
		if err != nil {
			return nil, pack.ReturnError(errno.DatabaseIfCommentExistError, err)
		}
		if !exist {
			return nil, pack.ReturnError(errno.CommentNotExistError, nil)
		}

		if err = i.repo.LikeComment(ctx, req.GetUID(), req.GetCommentID()); err != nil {
			return nil, pack.ReturnError(errno.DatabaseLikeCommentError, err)
		}
	case consts.LikeVideoActionKey:
		exist, err = i.repo.WhetherVideoLikeItemExist(ctx, req.GetUID(), req.GetVideoID())
		if err != nil {
			return nil, pack.ReturnError(errno.DatabaseWhetherVideoLikeItemExistError, err)
		}
		if exist {
			return nil, pack.ReturnError(errno.VideoAlreadyLikedError, nil)
		}

		exist, err = i.repo.WhetherVideoExist(ctx, req.GetVideoID())
		if err != nil {
			return nil, pack.ReturnError(errno.DatabaseIfVideoExistError, err)
		}
		if !exist {
			return nil, pack.ReturnError(errno.VideoNotExistError, nil)
		}
		if err = i.repo.LikeVideo(ctx, req.GetUID(), req.GetVideoID()); err != nil {
			return nil, pack.ReturnError(errno.DatabaseLikeVideoError, err)
		}
	}
	return nil, nil
}

func (i *InteractionCase) Dislike(ctx context.Context, req *interaction.DislikeReq) (r *interaction.DislikeResp, err error) {
	switch req.GetActionType() {
	case consts.DisLikeCommentActionKey:
		if err = i.repo.DislikeComment(ctx, req.GetUID(), req.GetCommentID()); err != nil {
			return nil, pack.ReturnError(errno.DatabaseDislikeCommentError, err)
		}
	case consts.DisLikeVideoActionKey:
		if err = i.repo.DislikeVideo(ctx, req.GetUID(), req.GetVideoID()); err != nil {
			return nil, pack.ReturnError(errno.DatabaseDislikeVideoError, err)
		}
	}
	return nil, nil
}

func (i *InteractionCase) LikeList(ctx context.Context, req *interaction.LikeListReq) (r *interaction.LikeListResp, err error) {
	videos, err := i.repo.LikeList(ctx, req.GetUID(), req.GetPageNum(), req.GetPageSize())
	if err != nil {
		return nil, pack.ReturnError(errno.DatabaseLikeListError, err)
	}

	result := make([]*video.VideoInfo, len(videos))
	for index := range videos {
		result[index] = repoV2IDL(&videos[index])
	}

	r = new(interaction.LikeListResp)
	count := int32(len(videos))
	r.SetCount(&count)
	r.SetItems(result)
	return r, nil
}

func (i *InteractionCase) Comment(ctx context.Context, req *interaction.CommentReq) (r *interaction.CommentResp, err error) {
	dto := &commentDTO{
		cid:     i.service.GenerateID(),
		uid:     req.GetUID(),
		vid:     req.GetVideoID(),
		content: req.GetContent(),
	}

	cid := i.service.GenerateID()
	if req.CommentID != nil { // 说明是对评论的评论
		dto.replyID = req.GetCommentID()
		rootId, err := i.repo.GetCommentRootID(ctx, dto.replyID)
		if err != nil {
			return nil, pack.ReturnError(errno.DatabaseGetCommentRootIDError, err)
		}
		if rootId == 0 {
			rootId = cid
		}
		dto.rootID = rootId
		if err = i.repo.CreateComment(ctx, dtoC2Repo(dto)); err != nil {
			return nil, pack.ReturnError(errno.DatabaseCreateCommentError, err)
		}
	} else { // 说明是对视频的评论
		dto.replyID = 0
		dto.cid = cid
		dto.rootID = 0
		if err = i.repo.CreateComment(ctx, dtoC2Repo(dto)); err != nil {
			return nil, pack.ReturnError(errno.DatabaseCreateCommentError, err)
		}
	}
	return nil, nil
}

func (i *InteractionCase) CommentList(ctx context.Context, req *interaction.CommentListReq) (r *interaction.CommentListResp, err error) {
	var comments []repository.Comment
	if req.VideoID != nil {
		comments, err = i.repo.GetVideoDirectCommentList(ctx, req.GetVideoID(), req.GetPageNum(), req.GetPageSize())
	} else if req.CommentID != nil {
		comments, err = i.repo.GetCommentList(ctx, req.GetCommentID(), req.GetPageNum(), req.GetPageSize())
	} else {
		return nil, err
	}
	if err != nil {
		return nil, pack.ReturnError(errno.DatabaseGetCommentListError, err)
	}

	commentInfos := make([]*interaction.CommentInfo, len(comments))
	for index := range comments {
		commentInfos[index] = repoC2IDL(&comments[index])
	}

	r = new(interaction.CommentListResp)
	length := int32(len(comments))
	r.SetCount(&length)
	r.SetItems(commentInfos)
	return r, nil
}

func (i *InteractionCase) DeleteComment(ctx context.Context, req *interaction.DeleteCommentReq) (r *interaction.DeleteCommentResp, err error) {
	if err = i.repo.DeleteComment(ctx, req.GetUID(), req.GetCommentID()); err != nil {
		return nil, pack.ReturnError(errno.DatabaseDeleteCommentError, err)
	}
	return nil, nil
}

func dtoC2Repo(dto *commentDTO) *repository.Comment {
	return &repository.Comment{
		ID:       dto.cid,
		UID:      dto.uid,
		VID:      dto.vid,
		RootID:   dto.rootID,
		ReplyID:  dto.replyID,
		Likes:    0,
		Content:  dto.content,
		CreateAt: time.Now().Unix(),
		DeleteAt: 0,
	}
}

func repoC2IDL(comment *repository.Comment) *interaction.CommentInfo {
	id := strconv.FormatInt(comment.ID, 10)
	vid := strconv.FormatInt(comment.VID, 10)
	uid := strconv.FormatInt(comment.UID, 10)
	replyID := strconv.FormatInt(comment.ReplyID, 10)
	rootID := strconv.FormatInt(comment.RootID, 10)
	createAt := time.Unix(comment.CreateAt, 0).Format("2006-01-02 15:04:05")
	return &interaction.CommentInfo{
		ID:       &id,
		UID:      &uid,
		VID:      &vid,
		RootID:   &rootID,
		ReplyID:  &replyID,
		Content:  &comment.Content,
		Likes:    &comment.Likes,
		CreateAt: &createAt,
	}
}
