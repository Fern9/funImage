package routers

import (
	"github.com/astaxie/beego"
	"github.com/zhusun/funImage/controllers/user"
)

func init() {
	ns := beego.NewNamespace(
		"/v1",
		beego.NSNamespace(
			"/auth",
			beego.NSRouter("/miniProgram/login", &user.LoginController{}, "post:MiniProgramLogin"),
		),
	)
	beego.AddNamespace(ns)
}
