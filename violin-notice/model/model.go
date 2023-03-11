package model

import "go.mongodb.org/mongo-driver/bson"

type TBlog struct {
	Bid   string `json:"bid"`
	Title string `json:"title"`
}

type TReminder struct {
	ReminderId string `bson:"reminder_id"`
	Title      string `bson:"title"`
	Info       string `bson:"info"`
	Time       string `bson:"time"`
	Type       bson.A `bson:"type"`
	TenantId   string `bson:"tenant_id"`
}
