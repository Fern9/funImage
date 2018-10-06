package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(
		"default",
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/funimage?charset=utf8mb4",
		30,
	)

}
