package main

import (
	"github.com/astaxie/beego"
	_ "github.com/zhusun/funImage/models"
	_ "github.com/zhusun/funImage/routers"
)

func main() {
	beego.Run()
}
