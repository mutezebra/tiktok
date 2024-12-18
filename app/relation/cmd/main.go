package main

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"

	"github.com/mutezebra/tiktok/pkg/trace"

	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"

	"github.com/mutezebra/tiktok/app/relation/cmd/pack"
	"github.com/mutezebra/tiktok/app/relation/config"
	"github.com/mutezebra/tiktok/app/relation/interface/persistence/database"
	"github.com/mutezebra/tiktok/pkg/consts"
	"github.com/mutezebra/tiktok/pkg/inject"
	"github.com/mutezebra/tiktok/pkg/kitex_gen/api/relation/relationservice"
	"github.com/mutezebra/tiktok/pkg/log"
)

func main() {
	RelationInit()
	registry := inject.NewRegistry(pack.RelationRegistry())
	defer registry.Close()
	registry.MustRegister(context.Background())

	addr, _ := net.ResolveTCPAddr("tcp", config.Conf.Service[consts.RelationServiceName].Address)
	srv := relationservice.NewServer(pack.ApplyRelation(),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.RelationServiceName}),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithMiddleware(trace.ServerTraceMiddleware(consts.RelationServiceName)),
	)

	if err := srv.Run(); err != nil {
		log.LogrusObj.Panic(err)
	}
}

func RelationInit() {
	config.InitConfig()
	log.InitLog(config.Conf.System.Status, config.Conf.System.OS)
	database.InitMysql()
}
