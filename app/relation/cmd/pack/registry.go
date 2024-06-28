package pack

import (
	relation "github.com/Mutezebra/tiktok/app/relation/domain/service"
	"github.com/Mutezebra/tiktok/app/relation/interface/persistence/database"
	"github.com/Mutezebra/tiktok/app/relation/usecase"
	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/pkg/inject"
)

func RelationRegistry() *inject.Registry {
	return &inject.Registry{
		Addr:   config.Conf.Service[consts.RelationServiceName].Address,
		Key:    config.Conf.Service[consts.RelationServiceName].ServiceName + "/" + config.Conf.Service[consts.RelationServiceName].Address,
		Prefix: config.Conf.Etcd.ServicePrefix}
}

func ApplyRelation() *usecase.RelationCase {
	repo := database.NewRelationRepository()
	service := relation.NewService(repo)
	return usecase.NewRelationCase(service, repo)
}
