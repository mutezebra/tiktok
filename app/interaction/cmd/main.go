package main

import (
	"context"
	"net"

	"github.com/Mutezebra/tiktok/app/interaction/cmd/pack"
	"github.com/Mutezebra/tiktok/app/interaction/interface/persistence/database"
	"github.com/Mutezebra/tiktok/pkg/inject"

	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"

	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/kitex_gen/api/interaction/interactionservice"
	"github.com/Mutezebra/tiktok/pkg/log"
)

func main() {
	InteractionInit()
	registry := inject.NewRegistry(pack.InteractionRegistry())
	defer registry.Close()
	registry.MustRegister(context.Background())

	addr, _ := net.ResolveTCPAddr("tcp", config.Conf.Service[consts.InteractionServiceName].Address)
	srv := interactionservice.NewServer(
		pack.ApplyInteraction(),
		server.WithServiceAddr(addr),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler))
	err := srv.Run()
	if err != nil {
		log.LogrusObj.Panic(err)
	}
}

func InteractionInit() {
	config.InitConfig()
	log.InitLog()
	database.InitMysql()
}
