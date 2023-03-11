package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	noticeServiceV1 "violin-home.cn/violin-notice/pkg/service/notice.service.v1"
	reminderServiceV1 "violin-home.cn/violin-notice/pkg/service/reminder.service.v1"
)

var Clinet noticeServiceV1.NoticeServiceClient
var Clinet2 reminderServiceV1.ReminderServiceClient

func InitGrpcClient() {
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("can not connect grpc server")
	}

	Clinet = noticeServiceV1.NewNoticeServiceClient(conn)
	Clinet2 = reminderServiceV1.NewReminderServiceClient(conn)

}
