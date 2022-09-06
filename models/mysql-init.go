package models

import (
	"github.com/astaxie/beego/config/env"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitORM() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	userName, _ := env.MustGet("DATABASE_USERNAME")
	password, _ := env.MustGet("DATABASE_PASSWORD")
	dbBaseUrl, _ := env.MustGet("DATABASE_URL")
	daPort, _ := env.MustGet("DATABASE_PORT")
	dbName, _ := env.MustGet("DATABASE_NAME")
	if registDataBaseErr := orm.RegisterDataBase("default", "mysql",
		userName+":"+password+"@tcp("+dbBaseUrl+":"+daPort+")/"+dbName+"?charset=utf8mb4&loc=Asia%2FShanghai"); registDataBaseErr != nil {
		logs.Error("[mysql]  注册连接数据库失败，失败信息：", registDataBaseErr.Error())
		panic(registDataBaseErr)
	} else {
		logs.Info("[mysql]  连接成功")
	}
}
