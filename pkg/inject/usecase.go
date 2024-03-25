package inject

import (
	"github.com/Mutezebra/tiktok/app/domain/service/user"
	"github.com/Mutezebra/tiktok/app/interface/persistence/database"
	"github.com/Mutezebra/tiktok/app/usecase"
	"github.com/Mutezebra/tiktok/pkg/oss"
)

func ApplyUser() *usecase.UserCase {
	repo := database.NewUserRepository()
	ossModel := oss.NewOssModel()
	service := user.NewService(repo, ossModel)
	return usecase.NewUseUseCase(repo, service)
}
