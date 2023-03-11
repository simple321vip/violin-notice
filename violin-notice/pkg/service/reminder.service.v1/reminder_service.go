package reminder_service_v1

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"violin-home.cn/common/logs"
	"violin-home.cn/violin-notice/config"
	"violin-home.cn/violin-notice/model"
)

type ReminderService struct {
	UnimplementedReminderServiceServer
}

func New() *ReminderService {
	return &ReminderService{}
}

func (*ReminderService) CreateReminder(_ context.Context, msg *ReminderMessage) (*ReminderResponse, error) {

	var t bson.A
	for _, k := range msg.Type {
		t = append(t, k)
	}

	_, err := config.MongoDBClient.Collection("t_reminder").InsertOne(context.TODO(), &model.TReminder{
		ReminderId: msg.ReminderId,
		Title:      msg.Title,
		Info:       msg.Info,
		Time:       msg.Time,
		Type:       t,
		TenantId:   msg.TenantId,
	})
	if err != nil {
		logs.LG.Error(err.Error())
	}

	return &ReminderResponse{}, err
}

func (*ReminderService) UpdateReminder(_ context.Context, msg *ReminderMessage) (*ReminderResponse, error) {

	var t bson.A
	for _, k := range msg.Type {
		t = append(t, k)
	}

	filter := bson.D{
		{"reminder_id", msg.ReminderId},
	}

	_, err := config.MongoDBClient.Collection("t_reminder").UpdateOne(context.TODO(), filter, &model.TReminder{
		ReminderId: msg.ReminderId,
		Title:      msg.Title,
		Info:       msg.Info,
		Time:       msg.Time,
		Type:       t,
		TenantId:   msg.TenantId,
	})
	if err != nil {
		logs.LG.Error(err.Error())
	}

	return &ReminderResponse{}, err
}

func (*ReminderService) DeleteReminder(_ context.Context, msg *ReminderMessage) (*ReminderResponse, error) {

	filter := bson.D{
		{"reminder_id", msg.ReminderId},
	}
	_, err := config.MongoDBClient.Collection("t_reminder").DeleteOne(context.TODO(), filter)
	if err != nil {
		logs.LG.Error(err.Error())
	}

	return &ReminderResponse{}, err
}

func (*ReminderService) SelectReminder(_ context.Context, msg *ReminderMessage) (*SelectReminderResponse, error) {

	filter := bson.D{
		{"reminder_id", msg.ReminderId},
	}

	cur, err := config.MongoDBClient.Collection("t_reminder").Find(context.TODO(), filter)
	if err != nil {
		logs.LG.Error(err.Error())
		return nil, err
	}

	var results []*ReminderMessage
	for cur.Next(context.TODO()) {
		var result ReminderMessage
		err := cur.Decode(&result)
		results = append(results, &result)
		if err != nil {
			logs.LG.Error(err.Error())
			return nil, err
		}
	}
	return &SelectReminderResponse{
		Response: results,
	}, err
}
