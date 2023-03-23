package reminder

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
	"violin-home.cn/common"
	"violin-home.cn/common/logs"
	"violin-home.cn/violin-api/grpc"
	reminderServiceV1 "violin-home.cn/violin-notice/pkg/service/reminder.service.v1"
)

type Handler struct {
}

// RequestReminder https://pkg.go.dev/github.com/go-playground/validator/v10#section-readme /**
type RequestReminder struct {
	ReminderId   string   `form:"reminder_id" xml:"reminder_id"`
	Title        string   `form:"title" xml:"title" binding:"required"`
	Info         string   `form:"info" xml:"info" binding:"required"`
	Type         []string `json:"type" xml:"type" binding:"required" validate:"gte=1"`
	ReminderDate string   `json:"reminder_date" xml:"reminder_date" binding:"required" validate:"datetime"`
}

func (nh *Handler) CreateReminder(ctx *gin.Context) {

	var rr RequestReminder
	if err := ctx.ShouldBindJSON(&rr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(rr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := &common.Result{}

	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := grpc.Clinet2.CreateReminder(c, &reminderServiceV1.ReminderMessage{
		Title:    rr.Title,
		Info:     rr.Info,
		Type:     []string{},
		Time:     "",
		TenantId: "xxxx",
	})

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, result.Fail(2001, "s"))
		return
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
		ctx.JSON(http.StatusOK, result.Fail(2001, "s"))
	}
	ctx.JSON(http.StatusOK, result.Success(resp))

}

func (nh *Handler) UpdateReminder(ctx *gin.Context) {

	var rr RequestReminder
	if err := ctx.ShouldBindJSON(&rr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(rr)

	result := &common.Result{}

	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := grpc.Clinet2.UpdateReminder(c, &reminderServiceV1.ReminderMessage{
		ReminderId: rr.ReminderId,
		Title:      rr.Title,
		Info:       rr.Info,
		Type:       rr.Type,
		Time:       "",
		TenantId:   ctx.GetString("tenantid"),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.Fail(2001, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(resp))
}

func (nh *Handler) QueryReminder(ctx *gin.Context) {

	tenantId := ctx.GetHeader("tenantid")
	result := &common.Result{}

	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := grpc.Clinet2.SelectReminder(c, &reminderServiceV1.ReminderMessage{
		TenantId: tenantId,
	})

	if err != nil {
		logs.LG.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, result.Fail(20001, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(resp))

}
