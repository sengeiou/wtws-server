package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"wtws-server/models"
	_ "wtws-server/routers"
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error("[server]  service start failed!!!! connect database failed!!!!!!")
		}
	}()
	models.InitORM()
	models.InitRedisClient()
	models.InitMongoDBClient()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error("[server]  service start failed!!!!")
		}
	}()
	orm.Debug = true
	beego.Run()
}
