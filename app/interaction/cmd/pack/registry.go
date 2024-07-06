package pack

import (
	"fmt"

	interaction "github.com/mutezebra/tiktok/interaction/domain/service"
	"github.com/mutezebra/tiktok/interaction/interface/persistence/database"
	"github.com/mutezebra/tiktok/pkg/snowflake"

	"github.com/mutezebra/tiktok/interaction/config"
	"github.com/mutezebra/tiktok/interaction/usecase"
	"github.com/mutezebra/tiktok/pkg/consts"
	"github.com/mutezebra/tiktok/pkg/inject"
)

func InteractionRegistry() *inject.Registry {
	id := snowflake.GenerateID(config.Conf.System.WorkerID, config.Conf.System.DataCenterID)
	return &inject.Registry{
		Addr:     config.Conf.Service[consts.InteractionServiceName].SvcAddress,
		Key:      config.Conf.Service[consts.InteractionServiceName].ServiceName + fmt.Sprintf("/%d", id),
		Prefix:   config.Conf.Etcd.ServicePrefix,
		EndPoint: config.Conf.Etcd.Endpoint,
	}
}

func ApplyInteraction() *usecase.InteractionCase {
	repo := database.NewInteractionRepository()
	service := &interaction.Service{
		Repo: repo,
	}
	return usecase.NewInteractionCase(repo, interaction.NewService(service))
}
