package pack

import (
	"fmt"

	"github.com/mutezebra/tiktok/pkg/consts"
	"github.com/mutezebra/tiktok/pkg/inject"
	"github.com/mutezebra/tiktok/pkg/oss"
	"github.com/mutezebra/tiktok/pkg/snowflake"
	"github.com/mutezebra/tiktok/video/config"
	video "github.com/mutezebra/tiktok/video/domain/service"
	"github.com/mutezebra/tiktok/video/interface/persistence/cache"
	"github.com/mutezebra/tiktok/video/interface/persistence/database"
	"github.com/mutezebra/tiktok/video/usecase"
)

func VideoRegistry() *inject.Registry {
	id := snowflake.GenerateID(config.Conf.System.WorkerID, config.Conf.System.DataCenterID)
	return &inject.Registry{
		Addr:     config.Conf.Service[consts.VideoServiceName].SvcAddress,
		Key:      config.Conf.Service[consts.VideoServiceName].ServiceName + fmt.Sprintf("/%d", id),
		Prefix:   config.Conf.Etcd.ServicePrefix,
		EndPoint: config.Conf.Etcd.Endpoint,
	}
}

func ApplyVideo() *usecase.VideoCase {
	repo := database.NewVideoRepository()
	ossModel := oss.NewOssModel(config.Conf.QiNiu.AvatarPath, config.Conf.QiNiu.VideoPath, config.Conf.QiNiu.CoverPath)
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
