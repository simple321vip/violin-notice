package v1

import (
	"github.com/gin-gonic/gin"
	"strings"
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
	rh := &reminder.Handler{}
	v1 := r.Group("violin-api/api/v1")
	{
		v1.Use(middle)
		v1.GET("/notice", sh.GetNotice)
		v1.POST("/reminder", rh.CreateReminder)
		v1.DELETE("/reminder", rh.DeleteReminder)
		v1.PUT("/reminder", rh.UpdateReminder)
		v1.GET("/reminder", rh.QueryReminder)
	}

}

func middle(c *gin.Context) {

	authorization := c.Param("authorization")
	strings.Split(authorization, ":")

	// redis
	//if json.User != "manu" || json.Password != "123" {
	//	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	//	return
	//}

	c.Next()
}
