package user

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	UserId   string `orm:"pk"`
	UserName string
	Mobile   string
	Email    string
	Avatar   string
}

func CreateUser(user User) bool {
	o := orm.NewOrm()
	// 查询用户是否存在
	u := User{UserId: user.UserId}
	err := o.Read(&u)
	if err == orm.ErrNoRows {
		o.Insert(&user)
		return true
	} else if err != nil {
		return false
	}
	o.Update(&user)
	return true
}
