package pack

import (
	user "github.com/Mutezebra/tiktok/app/user/domain/service"
	"github.com/Mutezebra/tiktok/app/user/interface/persistence/database"
	"github.com/Mutezebra/tiktok/app/user/usecase"
	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/pkg/inject"
	"github.com/Mutezebra/tiktok/pkg/oss"
	"github.com/Mutezebra/tiktok/pkg/utils"
)

func UserRegistry() *inject.Registry {
	return &inject.Registry{
		Addr:   config.Conf.Service[consts.UserServiceName].Address,
		Key:    config.Conf.Service[consts.UserServiceName].ServiceName + "/" + config.Conf.Service[consts.UserServiceName].Address,
		Prefix: config.Conf.Etcd.ServicePrefix}
}

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
