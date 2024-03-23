package pack

import (
	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/pkg/inject"
)

func UserRegistry() *inject.Registry {
	return &inject.Registry{
		Addr:   config.Conf.Service[consts.UserServiceName].Address,
		Key:    config.Conf.Service[consts.UserServiceName].ServiceName + "/" + config.Conf.Service[consts.UserServiceName].Address,
		Prefix: config.Conf.Etcd.ServicePrefix}
}
