package pack

import (
	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/pkg/inject"
)

func VideoRegistry() *inject.Registry {
	return &inject.Registry{
		Addr:   config.Conf.Service[consts.VideoServiceName].Address,
		Key:    config.Conf.Service[consts.VideoServiceName].ServiceName + "/" + config.Conf.Service[consts.VideoServiceName].Address,
		Prefix: config.Conf.Etcd.ServicePrefix}
}
