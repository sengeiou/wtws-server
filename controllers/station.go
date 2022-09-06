package controllers

import (
	"github.com/astaxie/beego"
	"wtws-server/common"
	"wtws-server/conf"
	"wtws-server/service"
)

// StationController User Api
type StationController struct {
	beego.Controller
}

// URLMapping url mapping
func (c *StationController) URLMapping() {
	c.Mapping("GetUserStation", c.GetUserStation)
}

// GetUserStation
// @Title GetUserStation
// @Description 获取用户所属服务站
// @router /list [get]
func (c *StationController) GetUserStation() {
	userInfo, _ := common.GetContextUserInfo(c.Ctx.Input.GetData(conf.CTX_CONTEXT_USER))

	c.Data["json"] = service.GetUserStation(userInfo.Id)
	c.ServeJSON()
}
