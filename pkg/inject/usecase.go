package inject

import (
	"github.com/Mutezebra/tiktok/app/domain/service/user"
	"github.com/Mutezebra/tiktok/app/interface/persistence/database"
	"github.com/Mutezebra/tiktok/app/usecase"
	"github.com/Mutezebra/tiktok/pkg/oss"
	"github.com/Mutezebra/tiktok/pkg/utils"
)

func ApplyUser() *usecase.UserCase {
	repo := database.NewUserRepository()
	ossModel := oss.NewOssModel()
	mfaModel := utils.NewMFAModel()
	service := user.NewService(repo, ossModel, mfaModel)
	return usecase.NewUseUseCase(repo, service)
}
