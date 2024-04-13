package rpc

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/streamclient"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"

	"github.com/Mutezebra/tiktok/app/domain/model"
	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/kitex_gen/api/interaction/interactionservice"
	"github.com/Mutezebra/tiktok/kitex_gen/api/relation/relationservice"
	"github.com/Mutezebra/tiktok/kitex_gen/api/user/userservice"
	"github.com/Mutezebra/tiktok/kitex_gen/api/video/videoservice"
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
	InteractionClient interactionservice.Client
	RelationClient    relationservice.Client
)

func Init() {
	Conf = config.Conf
	Resolver, err = inject.NewResolver()
	if err != nil {
		log.LogrusObj.Panic(err)
	}

	initClient(consts.UserServiceName)
	// initClient(consts.VideoServiceName)
	// initClient(consts.InteractionServiceName)
	//initClient(consts.RelationServiceName)
}

func initClient(serviceName string) {
	switch serviceName {
	case consts.UserServiceName:
		UserClient = userservice.MustNewClient(serviceName,
			client.WithHostPorts(discovery(serviceName)...),
			client.WithTransportProtocol(transport.TTHeader),
			client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	case consts.VideoServiceName:
		VideoClient = videoservice.MustNewClient(serviceName,
			client.WithHostPorts(discovery(serviceName)...),
			client.WithTransportProtocol(transport.TTHeader),
			client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
		VideoStreamClient = videoservice.MustNewStreamClient(serviceName,
			streamclient.WithHostPorts(discovery(serviceName)...),
			streamclient.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	case consts.InteractionServiceName:
		InteractionClient = interactionservice.MustNewClient(serviceName,
			client.WithHostPorts(discovery(serviceName)...),
			client.WithTransportProtocol(transport.TTHeader),
			client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	case consts.RelationServiceName:
		RelationClient = relationservice.MustNewClient(serviceName,
			client.WithHostPorts(discovery(serviceName)...),
			client.WithTransportProtocol(transport.TTHeader),
			client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		)
	}
}

func discovery(serviceName string) []string {
	prefix := Conf.Etcd.ServicePrefix + Conf.Service[serviceName].ServiceName
	addr, err := Resolver.ResolveWithPrefix(context.Background(), prefix)
	fmt.Println(addr)
	if err != nil {
		log.LogrusObj.Panic(err)
	}
	return addr
}
