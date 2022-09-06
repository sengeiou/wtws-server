package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config/env"
	"wtws-server/common"
	"wtws-server/controllers"
)

func init() {
	common.CorsDomain()

	version, _ := env.MustGet("VERSION")

	//token认证，并解析userID
	beego.InsertFilter(fmt.Sprintf("/%s/*", version), beego.BeforeRouter, common.Token)

	ns := beego.NewNamespace(fmt.Sprintf("/%s", version),
		beego.NSBefore(common.Auth),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/role",
			beego.NSInclude(
				&controllers.RoleController{},
			),
		),

		beego.NSNamespace("/station",
			beego.NSInclude(
				&controllers.StationController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
