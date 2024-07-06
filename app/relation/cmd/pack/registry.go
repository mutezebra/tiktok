package pack

import (
	"fmt"

	"github.com/mutezebra/tiktok/pkg/consts"
	"github.com/mutezebra/tiktok/pkg/inject"
	"github.com/mutezebra/tiktok/pkg/snowflake"
	"github.com/mutezebra/tiktok/relation/config"
	relation "github.com/mutezebra/tiktok/relation/domain/service"
	"github.com/mutezebra/tiktok/relation/interface/persistence/database"
	"github.com/mutezebra/tiktok/relation/usecase"
)

func RelationRegistry() *inject.Registry {
	id := snowflake.GenerateID(config.Conf.System.WorkerID, config.Conf.System.DataCenterID)
	return &inject.Registry{
		Addr:     config.Conf.Service[consts.RelationServiceName].SvcAddress,
		Key:      config.Conf.Service[consts.RelationServiceName].ServiceName + fmt.Sprintf("/%d", id),
		Prefix:   config.Conf.Etcd.ServicePrefix,
		EndPoint: config.Conf.Etcd.Endpoint,
	}
}

func ApplyRelation() *usecase.RelationCase {
	repo := database.NewRelationRepository()
	service := relation.NewService(repo)
	return usecase.NewRelationCase(service, repo)
}
