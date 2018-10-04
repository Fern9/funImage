package controllers

import (
	"github.com/astaxie/beego"
	"github.com/zhusun/funImage/pkg/e"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["json"] = e.Generate(e.NOT_AUTH, nil)
	c.ServeJSON()
}
