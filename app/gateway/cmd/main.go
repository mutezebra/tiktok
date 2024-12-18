package main

import (
	"os"
	"runtime/pprof"

	"github.com/mutezebra/tiktok/app/gateway/config"
	"github.com/mutezebra/tiktok/app/gateway/interface/persistence/database"
	"github.com/mutezebra/tiktok/app/gateway/interface/router"
	"github.com/mutezebra/tiktok/app/gateway/interface/rpc"
	"github.com/mutezebra/tiktok/pkg/log"
)

func main() {
	GatewayInit()

	h := router.NewRouter()

	f, err := os.Create("../datas/gateway-cpu.pprof")
	if err != nil {
		log.LogrusObj.Fatal(err)
	}
	if err = pprof.StartCPUProfile(f); err != nil {
		log.LogrusObj.Fatal(err)
	}
	defer pprof.StopCPUProfile()
	memFile, err := os.Create("../datas/gateway-mem.pprof")
	if err != nil {
		log.LogrusObj.Fatal(err)
	}
	if err = pprof.WriteHeapProfile(memFile); err != nil {
		log.LogrusObj.Fatal(err)
	}

	h.Spin()
}

func GatewayInit() {
	config.InitConfig()
	log.InitLog(config.Conf.System.Status, config.Conf.System.OS)
	database.InitMysql()
	// kafka.InitKafka(config.Conf.Kafka.Network, config.Conf.Kafka.Address)
	rpc.Init()
}
