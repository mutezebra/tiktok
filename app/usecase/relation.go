package usecase

import (
	"context"

	"github.com/Mutezebra/tiktok/app/domain/model/errno"
	"github.com/Mutezebra/tiktok/app/domain/repository"
	relationService "github.com/Mutezebra/tiktok/app/domain/service/relation"
	"github.com/Mutezebra/tiktok/app/usecase/pack"
	"github.com/Mutezebra/tiktok/kitex_gen/api/relation"
)

type RelationCase struct {
	service *relationService.Service
	repo    repository.RelationRepository
}

func NewRelationCase(service *relationService.Service, repo repository.RelationRepository) *RelationCase {
	return &RelationCase{service: service, repo: repo}
}

func (re *RelationCase) Follow(ctx context.Context, req *relation.FollowReq) (r *relation.FollowResp, err error) {
	if err = re.service.CheckUserExist(ctx, req.GetFollowerID()); err != nil {
		return nil, pack.ReturnError(errno.UserNotExistError, err)
	}

	var exist bool
	if exist, err = re.repo.WhetherFollowExist(ctx, req.GetUID(), req.GetFollowerID()); err != nil || exist {
		return nil, pack.ReturnError(errno.FollowAlreadyExistError, err)
	}

	if err = re.repo.Follow(ctx, req.GetUID(), req.GetFollowerID()); err != nil {
		return nil, pack.ReturnError(errno.DatabaseFollowError, err)
	}

	return nil, nil
}

func (re *RelationCase) GetFollowList(ctx context.Context, req *relation.GetFollowListReq) (r *relation.GetFollowListResp, err error) {
	var ids []int64
	if ids, err = re.repo.GetFollowList(ctx, req.GetUID()); err != nil {
		return nil, pack.ReturnError(errno.DatabaseGetFollowListError, err)
	}

	count := int32(len(ids))
	r = new(relation.GetFollowListResp)
	r.SetItems(ids)
	r.SetCount(&count)

	return r, nil
}

func (re *RelationCase) GetFansList(ctx context.Context, req *relation.GetFansListReq) (r *relation.GetFansListResp, err error) {
	var ids []int64
	if ids, err = re.repo.GetFansList(ctx, req.GetUID()); err != nil {
		return nil, pack.ReturnError(errno.DatabaseGetFansListError, err)
	}

	count := int32(len(ids))
	r = new(relation.GetFansListResp)
	r.SetItems(ids)
	r.SetCount(&count)

	return r, nil
}

func (re *RelationCase) GetFriendsList(ctx context.Context, req *relation.GetFriendsListReq) (r *relation.GetFriendsListResp, err error) {
	var ids []int64
	if ids, err = re.repo.GetFriendList(ctx, req.GetUID()); err != nil {
		return nil, pack.ReturnError(errno.DatabaseGetFriendsListError, err)
	}

	count := int32(len(ids))
	r = new(relation.GetFriendsListResp)
	r.SetItems(ids)
	r.SetCount(&count)

	return r, nil
}
