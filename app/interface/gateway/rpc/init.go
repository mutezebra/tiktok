package rpc

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"

	"github.com/Mutezebra/tiktok/app/domain/model"
	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/kitex_gen/api/user/userservice"
	"github.com/Mutezebra/tiktok/pkg/inject"
	"github.com/Mutezebra/tiktok/pkg/log"
)

var (
	err  error
	Conf *config.Config

	Resolver   model.Resolver
	UserClient userservice.Client
)

func Init() {
	Conf = config.Conf
	Resolver, err = inject.NewResolver()
	if err != nil {
		log.LogrusObj.Panic(err)
	}

	initClient(consts.UserServiceName)
}

func initClient(serviceName string) {
	switch serviceName {
	case consts.UserServiceName:
		UserClient, err = userservice.NewClient(serviceName, client.WithHostPorts(discovery(serviceName)...))
	}
	if err != nil {
		log.LogrusObj.Panic(err)
	}
}

func discovery(serviceName string) []string {
	addr := make([]string, 0)
	prefix := Conf.Etcd.ServicePrefix + Conf.Service[serviceName].ServiceName
	addr, err = Resolver.ResolveWithPrefix(context.Background(), prefix)
	fmt.Println(addr)
	if err != nil {
		log.LogrusObj.Panic(err)
	}
	return addr
}
