package pack

import (
	video "github.com/Mutezebra/tiktok/app/video/domain/service"
	"github.com/Mutezebra/tiktok/app/video/interface/persistence/cache"
	"github.com/Mutezebra/tiktok/app/video/interface/persistence/database"
	"github.com/Mutezebra/tiktok/app/video/usecase"
	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/pkg/inject"
	"github.com/Mutezebra/tiktok/pkg/oss"
)

func VideoRegistry() *inject.Registry {
	return &inject.Registry{
		Addr:   config.Conf.Service[consts.VideoServiceName].Address,
		Key:    config.Conf.Service[consts.VideoServiceName].ServiceName + "/" + config.Conf.Service[consts.VideoServiceName].Address,
		Prefix: config.Conf.Etcd.ServicePrefix}
}

func ApplyVideo() *usecase.VideoCase {
	repo := database.NewVideoRepository()
	ossModel := oss.NewOssModel()
	cacheRepo := cache.NewVideoCacheRepo()
	service := &video.Service{
		EnablePopularVideoRank:  true,
		EnableTimedRefreshViews: true,
		Repo:                    repo,
		Cache:                   cacheRepo,
		Oss:                     ossModel,
	}
	return usecase.NewVideoUseCase(repo, cacheRepo, video.NewService(service))
}
