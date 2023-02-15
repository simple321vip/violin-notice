package grpc

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"violin-home.cn/violin-notice/config"
	noticeServiceV1 "violin-home.cn/violin-notice/pkg/service/notice.service.v1"
)

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(server *grpc.Server)
}

func RegisterGrpcServer() *grpc.Server {

	c := &gRPCConfig{
		Addr: config.Conf.GC.Addr,
		RegisterFunc: func(server *grpc.Server) {
			noticeServiceV1.RegisterNoticeServiceServer(server, noticeServiceV1.New())
		},
	}

	s := grpc.NewServer()
	c.RegisterFunc(s)

	lis, err := net.Listen("tcp", c.Addr)

	if err != nil {
		log.Println("cannot listen")
	}

	go func() {
		err = s.Serve(lis)
		if err != nil {
			log.Println("Server started error")
		}

	}()
	log.Println("Server started！")
	return s

}
