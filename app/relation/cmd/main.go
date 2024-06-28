package main

import (
	"context"
	"net"

	"github.com/Mutezebra/tiktok/app/relation/interface/persistence/database"

	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"

	"github.com/Mutezebra/tiktok/app/relation/cmd/pack"
	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/kitex_gen/api/relation/relationservice"
	"github.com/Mutezebra/tiktok/pkg/inject"
	"github.com/Mutezebra/tiktok/pkg/log"
)

func main() {
	RelationInit()
	registry := inject.NewRegistry(pack.RelationRegistry())
	defer registry.Close()
	registry.MustRegister(context.Background())

	addr, _ := net.ResolveTCPAddr("tcp", config.Conf.Service[consts.RelationServiceName].Address)
	srv := relationservice.NewServer(
		pack.ApplyRelation(),
		server.WithServiceAddr(addr),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
	)
	err := srv.Run()
	if err != nil {
		log.LogrusObj.Panic(err)
	}
}

func RelationInit() {
	config.InitConfig()
	log.InitLog()
	database.InitMysql()
}
