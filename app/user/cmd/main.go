package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"runtime/pprof"

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

	go func() {
		// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.LogrusObj.Panic(err)
		}
		os.Exit(0)
	}()

	f, err := os.Create("../datas/user-cpu.pprof")
	if err != nil {
		log.LogrusObj.Fatal(err)
	}
	if err = pprof.StartCPUProfile(f); err != nil {
		log.LogrusObj.Fatal(err)
	}
	defer pprof.StopCPUProfile()
	memFile, err := os.Create("../datas/user-mem.pprof")
	if err != nil {
		log.LogrusObj.Fatal(err)
	}
	if err = pprof.WriteHeapProfile(memFile); err != nil {
		log.LogrusObj.Fatal(err)
	}

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
