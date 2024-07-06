package main

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"

	"github.com/mutezebra/tiktok/pkg/consts"
	"github.com/mutezebra/tiktok/pkg/inject"
	"github.com/mutezebra/tiktok/pkg/kitex_gen/api/user/userservice"
	"github.com/mutezebra/tiktok/pkg/log"
	"github.com/mutezebra/tiktok/pkg/oss"
	"github.com/mutezebra/tiktok/pkg/trace"
	"github.com/mutezebra/tiktok/user/cmd/pack"
	"github.com/mutezebra/tiktok/user/config"
	"github.com/mutezebra/tiktok/user/interface/persistence/database"
)

func main() {
	UserInit()
	registry := inject.NewRegistry(pack.UserRegistry())

	defer registry.Close()
	registry.MustRegister(context.Background())

	addr, err := net.ResolveTCPAddr("tcp", config.Conf.Service[consts.UserServiceName].Address)
	if err != nil {
		log.LogrusObj.Panic(err)
	}

	srv := userservice.NewServer(pack.ApplyUser(),
		server.WithServiceAddr(addr),
		server.WithMiddleware(trace.ServerTraceMiddleware("user")),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler))

	if err = srv.Run(); err != nil {
		log.LogrusObj.Panic(err)
	}
}

func UserInit() {
	config.InitConfig()
	log.InitLog(config.Conf.System.Status, config.Conf.System.OS)
	database.InitMysql()
	oss.InitOSS(config.Conf.QiNiu.AccessKey, config.Conf.QiNiu.SecretKey, config.Conf.QiNiu.Domain, config.Conf.QiNiu.Bucket)
}
