package pack

import (
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
