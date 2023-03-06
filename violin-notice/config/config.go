package config

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"violin-home.cn/common/logs"
)

var MongoDBClient *mongo.Database

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
	GC    *GrpcServerConfig
	LC    *logs.LogConfig
	MC    *MongoDBConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type GrpcServerConfig struct {
	Name string
	Addr string
}

type MongoDBConfig struct {
	Uri string
	DB  string
}

func InitConfig() *Config {
	v := viper.New()
	config := &Config{viper: v}
	workDir, _ := os.Getwd()
	if ginMODE := os.Getenv("GIN_MODE"); ginMODE == "release" {
		config.viper.SetConfigName("application_prod")
		config.viper.SetConfigType("yaml")
		config.viper.AddConfigPath(workDir + "/config")
	} else {
		config.viper.SetConfigName("application")
		config.viper.SetConfigType("yaml")
		config.viper.AddConfigPath(workDir + "/config")
	}
	err := v.ReadInConfig()

	if err != nil {
		log.Fatalln(err)
	}

	config.ReadServerConfig()
	config.ReadLogsConfig()
	config.ReadGrpcServerConfig()
	config.ReadMongoDBConfig()
	if err := logs.InitConfig(config.LC); err != nil {
		log.Println("************ log init failure **************")
	}

	logs.LG.Info("LOG INIT SUCCESSFUL")
	return config
}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{
		Name: c.viper.GetString("server.name"),
		Addr: c.viper.GetString("server.addr"),
	}
	c.SC = sc
}

func (c *Config) ReadGrpcServerConfig() {
	gsc := &GrpcServerConfig{
		Name: c.viper.GetString("grpc.name"),
		Addr: c.viper.GetString("grpc.addr"),
	}
	c.GC = gsc
}

func (c *Config) ReadLogsConfig() {

	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("zap.maxSize"),
		MaxAge:        c.viper.GetInt("zap.maxAge"),
		MaxBackups:    c.viper.GetInt("zap.maxBackups"),
	}

	c.LC = lc
}

func (c *Config) ReadMongoDBConfig() {

	c.MC = &MongoDBConfig{Uri: c.viper.GetString("mongodb.uri"), DB: c.viper.GetString("mongodb.db")}
}

func ConnectToDB(conf *Config) {
	clientOptions := options.Client().ApplyURI(conf.MC.Uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}
	logs.LG.Info("SUCCESSFULLY CONNECTED AND PINGED.")
	MongoDBClient = client.Database(conf.MC.DB)
}
