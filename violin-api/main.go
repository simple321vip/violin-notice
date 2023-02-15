package main

import (
	"github.com/gin-gonic/gin"
	"violin-home.cn/common"
	"violin-home.cn/violin-api/config"
	"violin-home.cn/violin-api/router"

	_ "violin-home.cn/violin-api/api"
)

func main() {

	// engine
	r := gin.Default()

	// router
	router.InitRouter(r)

	common.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, nil)

}
