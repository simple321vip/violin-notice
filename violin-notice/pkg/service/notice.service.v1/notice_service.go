package notice_service_v1

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/gomail.v2"
	"log"
	"violin-home.cn/common/logs"
	"violin-home.cn/violin-notice/config"
	"violin-home.cn/violin-notice/model"
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

	logs.LG.Debug("VIOLIN-NOTICE GRPC SERVICE CALLED SUCCESSFULLY")

	var TBlog *model.TBlog
	col := config.MongoDBClient.Collection("t_blog")
	filter := bson.D{{"bid", "2022103011580500000000000004"}}
	err := col.FindOne(context.TODO(), filter).Decode(&TBlog)
	if err != nil {
		logs.LG.Error("QUERY IS FAILURE.")
	}
	logs.LG.Debug(TBlog.Title)
	// 1. 获取参数

	// 2. 校验参数

	// 3. 生成验证码
	_ = &mail{
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
