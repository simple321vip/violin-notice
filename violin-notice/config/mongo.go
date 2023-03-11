package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"violin-home.cn/common/logs"
)

var MongoDBClient *mongo.Database

func ConnectToDB(conf *Config) {
	clientOptions := options.Client().ApplyURI(conf.MC.Uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}
	logs.LG.Info("MONGODB IS CONNECTED AND PINGED SUCCESSFULLY.")
	MongoDBClient = client.Database(conf.MC.DB)
}
