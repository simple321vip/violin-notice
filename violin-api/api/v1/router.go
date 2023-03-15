package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"violin-home.cn/violin-api/api/v1/notice"
	"violin-home.cn/violin-api/api/v1/reminder"
	"violin-home.cn/violin-api/redis"
	"violin-home.cn/violin-api/router"
)

func init() {
	router.Register(&Router{})
}

type Router struct {
}

func (sr *Router) Route(r *gin.Engine) {

	sh := &notice.Handler{}
	rh := &reminder.Handler{}
	v1 := r.Group("violin-api/api/v1", Interceptor())
	{
		v1.GET("/notice", sh.GetNotice)
		v1.POST("/reminder", rh.CreateReminder)
		v1.DELETE("/reminder", rh.DeleteReminder)
		v1.PUT("/reminder", rh.UpdateReminder)
		v1.GET("/reminder", rh.QueryReminder)
	}

}

func Interceptor() gin.HandlerFunc {

	return func(c *gin.Context) {
		tenantId := c.GetHeader("tenantid")

		if tenantId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "tenantid is empty"})
			c.Abort()
			return
		}

		authorization := c.GetHeader("authorization")
		var tmp []string
		if tmp = strings.Split(authorization, ":"); len(tmp) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			c.Abort()
			return
		}
		token := tmp[1]

		result, err := redis.ClientRedis.Get(tenantId).Result()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": err.Error()})
			c.Abort()
			return
		}

		if result != token {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "authorized error"})
			c.Abort()
			return
		}

		c.Next()
	}
}
