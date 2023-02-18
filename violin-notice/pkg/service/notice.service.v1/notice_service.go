package notice_service_v1

import (
	"context"
	"gopkg.in/gomail.v2"
	"log"
	"violin-home.cn/common/logs"
)

type NoticeService struct {
	UnimplementedNoticeServiceServer
}

func New() *NoticeService {
	return &NoticeService{}
}

type mail struct {
	senderAddr   string   // 发件人地址
	senderName   string   // 发件人名称
	receiverAddr []string // 收件人地址，可以有多个收件人
	subject      string   // 邮件主题
	text         string   // 正文
	host         string   // 邮件服务器地址
	port         int      // 邮件服务器端口号
	username     string   // 用户名
	password     string   // 密码或授权码
}

func (*NoticeService) SendNotice(
	ctx context.Context, msg *NoticeMessage) (*NoticeResponse, error) {

	logs.LG.Info("VIOLIN-NOTICE GRPC SERVICE CALLED SUCCESSFUL")
	// 1. 获取参数

	// 2. 校验参数

	// 3. 生成验证码
	m := &mail{
		senderAddr:   "simple321@vip.qq.com",
		senderName:   "guan",
		receiverAddr: nil,
		subject:      "",
		text:         "",
		host:         "",
		port:         0,
		username:     "",
		password:     "",
	}

	// 4. 调用短信平台
	logs.LG.Info(m.text)

	//SendMail(m)

	return &NoticeResponse{}, nil
}

func SendMail(s *mail) {
	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(s.senderAddr, s.senderName)}, // 发件人邮箱，发件人名称
		"To":      s.receiverAddr,                                // 多个收件人
		"Subject": {s.subject},                                   // 邮件主题
	})
	m.SetBody("text/plain", s.text)
	d := gomail.NewDialer(s.host, s.port, s.username, s.password) // 发送邮件服务器、端口号、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		log.Println("send mail err:", err)
	}
}
