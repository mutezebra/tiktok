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

	"github.com/mutezebra/tiktok/interaction/cmd/pack"
	"github.com/mutezebra/tiktok/interaction/config"
	"github.com/mutezebra/tiktok/interaction/interface/persistence/database"
	"github.com/mutezebra/tiktok/pkg/consts"
	"github.com/mutezebra/tiktok/pkg/inject"
	"github.com/mutezebra/tiktok/pkg/kitex_gen/api/interaction/interactionservice"
	"github.com/mutezebra/tiktok/pkg/log"
)

func main() {
	InteractionInit()
	registry := inject.NewRegistry(pack.InteractionRegistry())
	defer registry.Close()
	registry.MustRegister(context.Background())
	//tracer, closer, err := trace.NewTracer(consts.InteractionServiceName, config.Conf.Jaeger.CollectorEndpoint, config.Conf.Jaeger.AgentHostPort)
	//defer closer.Close()

	addr, _ := net.ResolveTCPAddr("tcp", config.Conf.Service[consts.InteractionServiceName].Address)
	srv := interactionservice.NewServer(pack.ApplyInteraction(),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.InteractionServiceName}),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithMiddleware(trace.ServerTraceMiddleware("interaction")),
	)

	f, err := os.Create("../datas/interaction-cpu.pprof")
	if err != nil {
		log.LogrusObj.Fatal(err)
	}
	if err = pprof.StartCPUProfile(f); err != nil {
		log.LogrusObj.Fatal(err)
	}
	defer pprof.StopCPUProfile()
	memFile, err := os.Create("../datas/interaction-mem.pprof")
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

func InteractionInit() {
	config.InitConfig()
	log.InitLog(config.Conf.System.Status, config.Conf.System.OS)
	database.InitMysql()
}
