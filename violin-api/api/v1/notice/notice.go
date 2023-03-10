package notice

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"violin-home.cn/common"
	"violin-home.cn/violin-api/grpc"
	noticeServiceV1 "violin-home.cn/violin-notice/pkg/service/notice.service.v1"
)

type Handler struct {
}

func (nh *Handler) GetNotice(ctx *gin.Context) {

	result := &common.Result{}
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := grpc.Clinet.SendNotice(c, &noticeServiceV1.NoticeMessage{})

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, result.Fail(2001, "s"))
	}
	log.Println(resp)
	ctx.JSON(http.StatusOK, result.Success("s"))

}
