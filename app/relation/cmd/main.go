package main

import (
	"context"
	"net"
	"os"
	"runtime/pprof"

	"github.com/cloudwego/kitex/pkg/rpcinfo"

	"github.com/mutezebra/tiktok/pkg/trace"

	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"

	"github.com/mutezebra/tiktok/pkg/consts"
	"github.com/mutezebra/tiktok/pkg/inject"
	"github.com/mutezebra/tiktok/pkg/kitex_gen/api/relation/relationservice"
	"github.com/mutezebra/tiktok/pkg/log"
	"github.com/mutezebra/tiktok/relation/cmd/pack"
	"github.com/mutezebra/tiktok/relation/config"
	"github.com/mutezebra/tiktok/relation/interface/persistence/database"
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

	f, err := os.Create("../datas/relation-cpu.pprof")
	if err != nil {
		log.LogrusObj.Fatal(err)
	}
	if err = pprof.StartCPUProfile(f); err != nil {
		log.LogrusObj.Fatal(err)
	}
	defer pprof.StopCPUProfile()
	memFile, err := os.Create("../datas/relation-mem.pprof")
	if err != nil {
		log.LogrusObj.Fatal(err)
	}
	if err = pprof.WriteHeapProfile(memFile); err != nil {
		log.LogrusObj.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.LogrusObj.Panic(err)
	}
}

func RelationInit() {
	config.InitConfig()
	log.InitLog(config.Conf.System.Status, config.Conf.System.OS)
	database.InitMysql()
}
