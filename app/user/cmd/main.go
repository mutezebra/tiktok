package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Mutezebra/tiktok/app/user/cmd/pack"
	"github.com/Mutezebra/tiktok/app/user/interface/persistence/database"

	"github.com/cloudwego/kitex/pkg/transmeta"

	"github.com/Mutezebra/tiktok/pkg/oss"

	"github.com/cloudwego/kitex/server"

	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/kitex_gen/api/user/userservice"
	"github.com/Mutezebra/tiktok/pkg/inject"
	"github.com/Mutezebra/tiktok/pkg/log"
)

func main() {
	UserInit()
	fmt.Println("have init")
	registry := inject.NewRegistry(pack.UserRegistry())
	defer registry.Close()
	registry.MustRegister(context.Background())

	addr, _ := net.ResolveTCPAddr("tcp", config.Conf.Service[consts.UserServiceName].Address)
	srv := userservice.NewServer(
		pack.ApplyUser(),
		server.WithServiceAddr(addr),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler))
	err := srv.Run()
	if err != nil {
		log.LogrusObj.Panic(err)
	}
}

func UserInit() {
	config.InitConfig()
	log.InitLog()
	database.InitMysql()
	oss.InitOSS()
}
