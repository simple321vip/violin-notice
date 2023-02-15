package main

import (
	"github.com/gin-gonic/gin"
	"violin-home.cn/common"
	"violin-home.cn/violin-notice/config"
	"violin-home.cn/violin-notice/grpc"
	"violin-home.cn/violin-notice/router"

	_ "violin-home.cn/violin-notice/api"
)

func main() {

	// engine
	r := gin.Default()

	// router
	router.InitRouter(r)

	gc := grpc.RegisterGrpcServer()
	stop := func() {
		gc.Stop()
	}
	common.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, stop)

}
