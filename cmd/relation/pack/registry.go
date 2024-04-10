package pack

import (
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
