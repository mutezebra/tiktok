package inject

import (
	"github.com/Mutezebra/tiktok/app/domain/service/interaction"
	"github.com/Mutezebra/tiktok/app/domain/service/user"
	"github.com/Mutezebra/tiktok/app/domain/service/video"
	"github.com/Mutezebra/tiktok/app/interface/persistence/cache"
	"github.com/Mutezebra/tiktok/app/interface/persistence/database"
	"github.com/Mutezebra/tiktok/app/usecase"
	"github.com/Mutezebra/tiktok/pkg/oss"
	"github.com/Mutezebra/tiktok/pkg/utils"
)

func ApplyUser() *usecase.UserCase {
	repo := database.NewUserRepository()
	ossModel := oss.NewOssModel()
	mfaModel := utils.NewMFAModel()

	service := &user.Service{
		Repo: repo,
		MFA:  mfaModel,
		OSS:  ossModel,
	}
	return usecase.NewUserUseCase(repo, user.NewService(service))
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

func ApplyInteraction() *usecase.InteractionCase {
	repo := database.NewInteractionRepository()
	service := &interaction.Service{
		Repo: repo,
	}
	return usecase.NewInteractionCase(repo, interaction.NewService(service))
}
