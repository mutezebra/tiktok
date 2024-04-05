package rpc

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client/streamclient"

	"github.com/Mutezebra/tiktok/kitex_gen/api/video/videoservice"

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

	Resolver          model.Resolver
	UserClient        userservice.Client
	VideoClient       videoservice.Client
	VideoStreamClient videoservice.StreamClient
)

func Init() {
	Conf = config.Conf
	Resolver, err = inject.NewResolver()
	if err != nil {
		log.LogrusObj.Panic(err)
	}

	initClient(consts.UserServiceName)
	initClient(consts.VideoServiceName)
}

func initClient(serviceName string) {
	switch serviceName {
	case consts.UserServiceName:
		UserClient = userservice.MustNewClient(serviceName, client.WithHostPorts(discovery(serviceName)...))
	case consts.VideoServiceName:
		VideoClient = videoservice.MustNewClient(serviceName, client.WithHostPorts(discovery(serviceName)...))
		VideoStreamClient = videoservice.MustNewStreamClient(serviceName, streamclient.WithHostPorts(discovery(serviceName)...))
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
