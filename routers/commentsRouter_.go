package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["wtws-server/controllers:RoleController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:RoleController"],
        beego.ControllerComments{
            Method: "DeleteRole",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:RoleController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:RoleController"],
        beego.ControllerComments{
            Method: "AddRole",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:RoleController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:RoleController"],
        beego.ControllerComments{
            Method: "GetAllRoleList",
            Router: "/all-list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:RoleController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:RoleController"],
        beego.ControllerComments{
            Method: "GetRoleFunctions",
            Router: "/functions",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:RoleController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:RoleController"],
        beego.ControllerComments{
            Method: "AddRoleFunctions",
            Router: "/functions",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:StationController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:StationController"],
        beego.ControllerComments{
            Method: "GetUserStation",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddUserInfo",
            Router: "/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserLIst",
            Router: "/list",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUserLoginName",
            Router: "/login-name",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "LogOut",
            Router: "/logout",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "RestUserPwd",
            Router: "/reset-pwd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdatePwd",
            Router: "/update-pwd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserInfo",
            Router: "/user-info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
