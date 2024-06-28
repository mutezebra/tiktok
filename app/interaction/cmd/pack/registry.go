package pack

import (
	interaction "github.com/Mutezebra/tiktok/app/interaction/domain/service"
	"github.com/Mutezebra/tiktok/app/interaction/interface/persistence/database"

	"github.com/Mutezebra/tiktok/app/interaction/usecase"
	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/pkg/inject"
)

func InteractionRegistry() *inject.Registry {
	return &inject.Registry{
		Addr:   config.Conf.Service[consts.InteractionServiceName].Address,
		Key:    config.Conf.Service[consts.InteractionServiceName].ServiceName + "/" + config.Conf.Service[consts.InteractionServiceName].Address,
		Prefix: config.Conf.Etcd.ServicePrefix}
}

func ApplyInteraction() *usecase.InteractionCase {
	repo := database.NewInteractionRepository()
	service := &interaction.Service{
		Repo: repo,
	}
	return usecase.NewInteractionCase(repo, interaction.NewService(service))
}
