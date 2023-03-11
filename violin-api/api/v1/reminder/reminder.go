package reminder

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"violin-home.cn/common"
	"violin-home.cn/violin-api/grpc"
	reminderServiceV1 "violin-home.cn/violin-notice/pkg/service/reminder.service.v1"
)

type Handler struct {
}

func (nh *Handler) CreateReminder(ctx *gin.Context) {

	result := &common.Result{}

	title := ctx.PostForm("title")
	info := ctx.PostForm("info")
	//t := ctx.PostForm("type")
	ti := ctx.PostForm("time")

	tenantId := ctx.PostForm("time")

	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := grpc.Clinet2.CreateReminder(c, &reminderServiceV1.ReminderMessage{
		ReminderId: "3333333",
		Title:      title,
		Info:       info,
		Type:       nil,
		Time:       ti,
		TenantId:   tenantId,
	})

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, result.Fail(2001, "s"))
	}
	log.Println(resp)
	ctx.JSON(http.StatusOK, result.Success(resp))

}

func (nh *Handler) DeleteReminder(ctx *gin.Context) {

	reminderId := ctx.PostForm("reminderId")

	result := &common.Result{}
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := grpc.Clinet2.DeleteReminder(c, &reminderServiceV1.ReminderMessage{
		ReminderId: reminderId,
	})

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, result.Fail(2001, "s"))
	}
	log.Println(resp)
	ctx.JSON(http.StatusOK, result.Success(resp))

}

func (nh *Handler) UpdateReminder(ctx *gin.Context) {

	result := &common.Result{}

	reminderId := ctx.PostForm("reminderId")
	title := ctx.PostForm("title")
	info := ctx.PostForm("info")
	//t := ctx.PostForm("type")
	ti := ctx.PostForm("time")

	tenantId := ctx.PostForm("time")

	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := grpc.Clinet2.UpdateReminder(c, &reminderServiceV1.ReminderMessage{
		ReminderId: reminderId,
		Title:      title,
		Info:       info,
		Type:       []string{"mail", "phone"},
		Time:       ti,
		TenantId:   tenantId,
	})

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, result.Fail(2001, "s"))
	}
	log.Println(resp)
	ctx.JSON(http.StatusOK, result.Success(resp))

}

func (nh *Handler) QueryReminder(ctx *gin.Context) {

	result := &common.Result{}
	reminderId := ctx.PostForm("reminderId")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := grpc.Clinet2.SelectReminder(c, &reminderServiceV1.ReminderMessage{
		ReminderId: reminderId,
	})

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, result.Fail(2001, "s"))
	}
	log.Println(resp)
	ctx.JSON(http.StatusOK, result.Success(resp))

}
