package models

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/config/env"
	"github.com/astaxie/beego/logs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongodbClientDB *mongo.Database

//InitMongoDBClient 初始化mongodb连接
func InitMongoDBClient() {
	mongodbUri, _ := env.MustGet("MONGODB_URI")
	mongodbSource, _ := env.MustGet("MONGODB_SOURCE")

	// Set client options 设置连接参数
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("%s/?connect=direct;authSource=%s", mongodbUri, mongodbSource))

	// Connect to MongoDB 连接数据库
	if client, err := mongo.Connect(context.TODO(), clientOptions); err != nil {
		logs.Error("[mongodb] mongodb 连接失败，创建client失败", err)
		panic(err)
	} else if err = client.Ping(context.TODO(), nil); err != nil {
		// Check the connection 测试连接
		logs.Error("[mongodb] mongodb 连接失败，Ping失败", err)
		panic(err)
	} else {
		logs.Info("[mongodb] mongodb 连接成功")
		MongodbClientDB = client.Database(mongodbSource)
	}

}
