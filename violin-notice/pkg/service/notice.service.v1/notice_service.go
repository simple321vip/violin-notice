package notice_service_v1

import (
	"context"
	"log"
)

type NoticeService struct {
	UnimplementedNoticeServiceServer
}

func New() *NoticeService {
	return &NoticeService{}
}

func (*NoticeService) SendNotice(ctx context.Context, msg *NoticeMessage) (*NoticeResponse, error) {

	// 1. 获取参数

	// 2. 校验参数

	// 3. 生成验证码

	// 4. 调用短信平台
	log.Println("grpc is called")

	return &NoticeResponse{}, nil
}
