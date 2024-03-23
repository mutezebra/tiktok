package main

import (
	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/pkg/log"
)

func main() {
	config.InitConfig()
	log.InitLog()
	log.LogrusObj.Panic("aaa panic")

}
