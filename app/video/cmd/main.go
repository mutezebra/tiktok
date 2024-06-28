package main

import (
	"context"
	"net"

	"github.com/Mutezebra/tiktok/app/video/cmd/pack"
	"github.com/Mutezebra/tiktok/app/video/interface/persistence/cache"
	"github.com/Mutezebra/tiktok/app/video/interface/persistence/database"

	"github.com/cloudwego/kitex/pkg/transmeta"

	"github.com/Mutezebra/tiktok/kitex_gen/api/video/videoservice"

	"github.com/Mutezebra/tiktok/pkg/oss"

	"github.com/cloudwego/kitex/server"

	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/pkg/inject"
	"github.com/Mutezebra/tiktok/pkg/log"
)

func main() {
	VideoInit()
	registry := inject.NewRegistry(pack.VideoRegistry())
	defer registry.Close()
	registry.MustRegister(context.Background())

	addr, _ := net.ResolveTCPAddr("tcp", config.Conf.Service[consts.VideoServiceName].Address)
	srv := videoservice.NewServer(
		pack.ApplyVideo(),
		server.WithServiceAddr(addr),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler))
	err := srv.Run()
	if err != nil {
		log.LogrusObj.Panic(err)
	}
}

func VideoInit() {
	config.InitConfig()
	log.InitLog()
	database.InitMysql()
	cache.InitCache()
	oss.InitOSS()
}
