package user

import (
	"github.com/astaxie/beego"
	"github.com/zhusun/funImage/pkg/e"
	userCache "github.com/zhusun/funImage/pkg/cache/user"
	"github.com/zhusun/funImage/models/user"
)

type UserController struct {
	beego.Controller
}

func (c *UserController)GetUserInfo() {
	token := c.GetString("token")
	if token == "" {
		c.Data["json"] = e.Generate(e.ParamsError, "")
		c.ServeJSON()
	}
	userInfo := userCache.Get(token)
	if userInfo == nil {
		c.Data["json"] = e.Generate(e.GetUserInfoError, nil)
		c.ServeJSON()
	}
	userInfo = userInfo.(user.User)
	c.Data["json"] = e.Generate(e.Success, userInfo)
	c.ServeJSON()
}

