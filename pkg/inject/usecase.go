package inject

import (
	"github.com/Mutezebra/tiktok/app/domain/service/user"
	"github.com/Mutezebra/tiktok/app/interface/persistence/database"
	"github.com/Mutezebra/tiktok/app/usecase"
)

func ApplyUser() *usecase.UserCase {
	repo := database.NewUserRepository()
	service := user.NewService(repo)
	return usecase.NewUseUseCase(repo, service)
}
