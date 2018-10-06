package user

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/zhusun/funImage/models/user"
	"time"
)

var userCache, _ = cache.NewCache("memory", `{"interval":60}`)

func Put(token string, user user.User) error {
	expireTime, err := beego.AppConfig.Int64("userTokenExpire")
	if err != nil {
		return err
	}
	err = userCache.Put(token, user, time.Duration(expireTime))
	if err != nil {
		return err
	}
	return nil
}

func Get(token string) (user.User) {
	u := userCache.Get(token)
	return u.(user.User)
}

func Delete(token string) error {
	err := userCache.Delete(token)
	return err
}
