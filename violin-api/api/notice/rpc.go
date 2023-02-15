package notice

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	noticeServiceV1 "violin-home.cn/violin-notice/pkg/service/notice.service.v1"
)

var Clinet noticeServiceV1.NoticeServiceClient

func InitGrpcClient() {
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("can not connect grpc server")
	}

	Clinet = noticeServiceV1.NewNoticeServiceClient(conn)

}
