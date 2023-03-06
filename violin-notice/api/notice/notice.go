package notice

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	clientV3 "go.etcd.io/etcd/client/v3"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
	"violin-home.cn/common"
	"violin-home.cn/violin-notice/config"
)

type Handler struct {
}

func (*Handler) getNotice(ctx *gin.Context) {

	col := config.MongoDBClient.Collection("t_blog")
	filter := bson.D{{"bid", "2022103011580500000000000004"}}
	re := col.FindOne(context.TODO(), filter)

	fmt.Println(re)

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

func SendNotice() error {
	cli, err := clientV3.New(clientV3.Config{Endpoints: []string{"0.0.0.0:2379"}, DialTimeout: 5 * time.Second, Username: "root", Password: "123654"})
	if err != nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err = cli.KV.Put(ctx, "111", "222")

	if err != nil {
		return nil
	}
	return err
}
