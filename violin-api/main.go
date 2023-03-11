package main

import (
	"github.com/gin-gonic/gin"
	"violin-home.cn/common"
	"violin-home.cn/violin-api/config"
	"violin-home.cn/violin-api/grpc"
	"violin-home.cn/violin-api/router"

	_ "violin-home.cn/violin-api/api/v1"
)

func main() {

	// engine
	r := gin.Default()

	// router
	router.InitRouter(r)

	// grpc client
	grpc.InitGrpcClient()

	common.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, nil)

}
