package notice

import (
	"github.com/gin-gonic/gin"
	"violin-home.cn/violin-notice/router"
)

func init() {
	router.Register(&Router{})
}

type Router struct {
}

func (sr *Router) Route(r *gin.Engine) {
	sh := &Handler{}
	r.GET("/violin-api/v1/settings", sh.getNotice)
}
