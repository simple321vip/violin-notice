package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"violin-home.cn/violin-notice/config"
	"violin-home.cn/violin-notice/grpc"
)

func main() {

	conf := config.InitConfig()

	grpc.RegisterGrpcServer(conf)

	config.ConnectToDB(conf)

	// 创建监听退出chan
	quit := make(chan os.Signal)

	// 监听指定信号 ctrl+c kill
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	s := <-quit
	fmt.Println(s)
}
