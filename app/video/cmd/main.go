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
	"github.com/mutezebra/tiktok/pkg/kitex_gen/api/video/videoservice"
	"github.com/mutezebra/tiktok/pkg/log"
	"github.com/mutezebra/tiktok/pkg/oss"
	"github.com/mutezebra/tiktok/video/cmd/pack"
	"github.com/mutezebra/tiktok/video/config"
	"github.com/mutezebra/tiktok/video/interface/persistence/cache"
	"github.com/mutezebra/tiktok/video/interface/persistence/database"
)

func main() {
	VideoInit()
	registry := inject.NewRegistry(pack.VideoRegistry())
	defer registry.Close()
	registry.MustRegister(context.Background())
	//tracer, closer, err := trace.NewTracer(consts.VideoServiceName, config.Conf.Jaeger.CollectorEndpoint, config.Conf.Jaeger.AgentHostPort)
	//defer closer.Close()

	addr, _ := net.ResolveTCPAddr("tcp", config.Conf.Service[consts.VideoServiceName].Address)
	srv := videoservice.NewServer(pack.ApplyVideo(),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServiceName}),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithMiddleware(trace.ServerTraceMiddleware(consts.VideoServiceName)),
	)

	f, err := os.Create("../datas/video-cpu.pprof")
	if err != nil {
		log.LogrusObj.Fatal(err)
	}
	if err = pprof.StartCPUProfile(f); err != nil {
		log.LogrusObj.Fatal(err)
	}
	defer pprof.StopCPUProfile()
	memFile, err := os.Create("../datas/video-mem.pprof")
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

func VideoInit() {
	config.InitConfig()
	log.InitLog(config.Conf.System.Status, config.Conf.System.OS)
	database.InitMysql()
	cache.InitCache()
	oss.InitOSS(config.Conf.QiNiu.AccessKey, config.Conf.QiNiu.SecretKey, config.Conf.QiNiu.Domain, config.Conf.QiNiu.Bucket)
}
