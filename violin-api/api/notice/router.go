package notice

import (
	"github.com/gin-gonic/gin"
	"violin-home.cn/violin-api/router"
)

func init() {
	router.Register(&Router{})
}

type Router struct {
}

func (sr *Router) Route(r *gin.Engine) {
	InitGrpcClient()
	sh := &Handler{}
	r.GET("violin-api/v1/notice", sh.getNotice)
}
