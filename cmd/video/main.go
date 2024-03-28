package main

import (
	"context"
	"github.com/Mutezebra/tiktok/cmd/video/pack"
	"github.com/Mutezebra/tiktok/kitex_gen/api/video/videoservice"
	"net"

	"github.com/Mutezebra/tiktok/pkg/oss"

	"github.com/Mutezebra/tiktok/app/interface/persistence/database"

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
	srv := videoservice.NewServer(inject.ApplyVideo(), server.WithServiceAddr(addr))
	err := srv.Run()
	if err != nil {
		log.LogrusObj.Panic(err)
	}
}

func VideoInit() {
	config.InitConfig()
	log.InitLog()
	database.InitMysql()
	oss.InitOSS()
}
