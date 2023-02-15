package notice

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"violin-home.cn/common"
)

type Handler struct {
}

func (*Handler) getNotice(ctx *gin.Context) {

	resp := &common.Result{}
	tenantId := ctx.Query("tenant_id")

	if !common.VerifyTenantID(tenantId) {
		ctx.JSON(http.StatusOK, resp.Fail(2001, "error"))
		return
	}

	go func() {
		time.Sleep(time.Second)
		log.Println("get notice success")
	}()

	ctx.JSON(http.StatusOK, resp.Success("s"))
}
