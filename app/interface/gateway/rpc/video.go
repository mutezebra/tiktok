package rpc

import (
	"context"
	"io"

	idl "github.com/Mutezebra/tiktok/kitex_gen/api/video"
	"github.com/Mutezebra/tiktok/pkg/log"
)

func VideoFeed(ctx context.Context, req *idl.VideoFeedReq, ch chan []byte) (err error) {
	stream, err := VideoStreamClient.VideoFeed(ctx, req)
	if err != nil {
		return err
	}
	var i int
	resp := new(idl.VideoFeedResp)
	for {
		resp, err = stream.Recv()
		if err == io.EOF {
			err = nil
			log.LogrusObj.Info("client stream closed count is", i)
			break
		} else if err != nil {
			log.LogrusObj.Info("err != nil", err)
			break
		}
		ch <- resp.Video
		i++
	}

	log.LogrusObj.Info("i am free")
	ch <- nil
	return err
}

func PublishVideo(ctx context.Context, req *idl.PublishVideoReq) (r *idl.PublishVideoResp, err error) {
	r, err = VideoClient.PublishVideo(ctx, req)
	return r, err
}

func GetVideoList(ctx context.Context, req *idl.GetVideoListReq) (r *idl.GetVideoListResp, err error) {
	r, err = VideoClient.GetVideoList(ctx, req)
	return r, err
}

func GetVideoPopular(ctx context.Context, req *idl.GetVideoPopularReq) (r *idl.GetVideoPopularResp, err error) {
	r, err = VideoClient.GetVideoPopular(ctx, req)
	return r, err
}

func SearchVideo(ctx context.Context, req *idl.SearchVideoReq) (r *idl.SearchVideoResp, err error) {
	r, err = VideoClient.SearchVideo(ctx, req)
	return r, err
}
