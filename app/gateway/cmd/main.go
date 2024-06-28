package main

import (
	"github.com/Mutezebra/tiktok/app/gateway/interface/persistence/database"
	"github.com/Mutezebra/tiktok/app/gateway/interface/router"
	"github.com/Mutezebra/tiktok/app/gateway/interface/rpc"
	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/pkg/kafka"
	"github.com/Mutezebra/tiktok/pkg/log"
)

func main() {
	GatewayInit()
	h := router.NewRouter()
	h.Spin()
}

func GatewayInit() {
	config.InitConfig()
	log.InitLog()
	database.InitMysql()
	kafka.InitKafka()
	rpc.Init()
}
