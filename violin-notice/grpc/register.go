package grpc

import (
	"google.golang.org/grpc"
	"net"
	"violin-home.cn/common/logs"
	"violin-home.cn/violin-notice/config"
	noticeServiceV1 "violin-home.cn/violin-notice/pkg/service/notice.service.v1"
)

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(server *grpc.Server)
}

func RegisterGrpcServer(conf *config.Config) *grpc.Server {

	c := &gRPCConfig{
		Addr: conf.GC.Addr,
		RegisterFunc: func(server *grpc.Server) {
			noticeServiceV1.RegisterNoticeServiceServer(server, noticeServiceV1.New())
		},
	}

	s := grpc.NewServer()
	c.RegisterFunc(s)

	lis, err := net.Listen("tcp", c.Addr)

	if err != nil {
		logs.LG.Error("CAN NOT TO LISTEN TO TCP ADDRESS [" + c.Addr + "] SUCCESSFULLY.")
	}
	logs.LG.Info("LISTEN TO TCP ADDRESS [" + c.Addr + "] SUCCESSFULLY.")

	go func() {
		err = s.Serve(lis)
		if err != nil {
			logs.LG.Error("GRPC SERVER STARTED ERROR.")
		}
		logs.LG.Info("SERVER STARTED SUCCESSFULLY.")
	}()

	return s

}
