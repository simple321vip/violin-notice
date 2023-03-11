package v1

import (
	"github.com/gin-gonic/gin"
	"violin-home.cn/violin-api/api/v1/notice"
	"violin-home.cn/violin-api/api/v1/reminder"
	"violin-home.cn/violin-api/router"
)

func init() {
	router.Register(&Router{})
}

type Router struct {
}

func (sr *Router) Route(r *gin.Engine) {

	sh := &notice.Handler{}
	r.GET("violin-api/v1/notice", sh.GetNotice)

	rh := &reminder.Handler{}
	r.GET("violin-api/v1/reminder", rh.CreateReminder)
}
