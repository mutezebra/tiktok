package usecase

import (
	"context"
	"fmt"
	"github.com/Mutezebra/tiktok/app/domain/model/errno"
	"github.com/Mutezebra/tiktok/app/domain/repository"
	videoService "github.com/Mutezebra/tiktok/app/domain/service/video"
	"github.com/Mutezebra/tiktok/app/usecase/pack"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/kitex_gen/api/video"
	"path/filepath"
	"strconv"
)

type VideoCase struct {
	repo    repository.VideoRepository
	cache   repository.VideoCacheRepository
	service *videoService.Service
}

func NewVideoUseCase(repo repository.VideoRepository, cache repository.VideoCacheRepository, service *videoService.Service) *VideoCase {
	return &VideoCase{
		repo:    repo,
		cache:   cache,
		service: service,
	}
}

type dtoVideo struct {
	vid      int64
	uid      int64
	videoURL string
	coverURL string
	intro    string
	title    string
}

func (v *VideoCase) VideoFeed(req *video.VideoFeedReq, stream video.VideoService_VideoFeedServer) (err error) {
	videoURL, err := v.repo.GetVideoUrl(context.Background(), req.GetVID())
	if err != nil {
		return pack.ReturnError(errno.DatabaseGetVideoUrlError, err)
	}

	data, err := v.service.VideoFeed(videoURL)
	if err != nil {
		return pack.ReturnError(errno.OssGetVideoFeedError, err)
	}

	resp := new(video.VideoFeedResp)
	begin := 0
	offset := 2 * consts.MB
	remain := len(data)
	for remain > 0 {
		if remain < offset {
			resp.Video = data[begin:]
		} else {
			resp.Video = data[begin : begin+offset]
			begin = begin + offset
		}
		remain = remain - offset
		err = stream.Send(resp)
		if err != nil {
			return pack.ReturnError(errno.VideoFeedStreamSendError, err)
		}
	}

	return nil
}

func (v *VideoCase) PublishVideo(ctx context.Context, req *video.PublishVideoReq) (r *video.PublishVideoResp, err error) {
	vid := v.service.GenerateID()
	videoName := fmt.Sprintf("%d%s", vid, filepath.Ext(req.GetVideoName()))
	err, videoURL := v.service.UploadVideo(ctx, videoName, req.GetVideo())
	if err != nil {
		return nil, pack.ReturnError(errno.OssUploadVideoCoverError, err)
	}

	coverName := fmt.Sprintf("%d%s", vid, filepath.Ext(req.GetCoverName()))
	err, coverURL := v.service.UploadVideoCover(ctx, coverName, req.GetCover())
	if err != nil {
		return nil, pack.ReturnError(errno.OssUploadVideoCoverError, err)
	}

	dto := &dtoVideo{
		vid:      vid,
		uid:      req.GetUID(),
		videoURL: videoURL,
		coverURL: coverURL,
		intro:    req.GetIntro(),
		title:    req.GetTitle(),
	}
	vid, err = v.repo.CreateVideo(ctx, dtoV2Repo(dto))
	if err != nil {
		return nil, pack.ReturnError(errno.DatabaseCreateVideoError, err)
	}

	r = new(video.PublishVideoResp)
	strVid := strconv.FormatInt(vid, 10)
	r.SetVID(&strVid)

	return r, nil
}

func (v *VideoCase) GetVideoList(ctx context.Context, req *video.GetVideoListReq) (r *video.GetVideoListResp, err error) {
	return nil, err

}

func (v *VideoCase) GetVideoPopular(ctx context.Context, req *video.GetVideoPopularReq) (r *video.GetVideoPopularResp, err error) {
	return nil, err

}

func (v *VideoCase) SearchVideo(ctx context.Context, req *video.SearchVideoReq) (r *video.SearchVideoResp, err error) {
	return nil, err

}

func dtoV2Repo(dto *dtoVideo) *repository.Video {
	return &repository.Video{
		ID:       dto.vid,
		UID:      dto.uid,
		VideoURL: dto.videoURL,
		CoverURL: dto.coverURL,
		Intro:    dto.intro,
		Title:    dto.title,
	}
}
