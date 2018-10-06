package user

import (
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"github.com/medivhzhan/weapp"

	"github.com/zhusun/funImage/models/user"
	"github.com/zhusun/funImage/pkg/e"
	userCache "github.com/zhusun/funImage/pkg/cache/user"
)

type LoginController struct {
	beego.Controller
}

func (login *LoginController) MiniProgramLogin() {
	rawData := login.GetString("rawData")
	encryptedData := login.GetString("encryptedData")
	signature := login.GetString("signature")
	iv := login.GetString("signature")
	ssk := beego.AppConfig.String("miniProgramKey")

	// 解密微信小程序用户信息
	userInfo, err := weapp.DecryptUserInfo(rawData, encryptedData, signature, iv, ssk)
	if err != nil {
		login.Data["json"] = e.Generate(e.GetMiniProgramUserInfoError, nil)
		login.ServeJSON()
	}

	//实例化用户对象
	u := user.User{
		UserId:   userInfo.OpenID,
		UserName: userInfo.Nickname,
		Mobile:   "",
		Email:    "",
		Avatar:   userInfo.Avatar,
	}

	//创建用户，如果已存在，则更新，不存在则创建
	result := user.CreateUser(u)
	if result == false {
		login.Data["json"] = e.Generate(e.AuthError, "")
		login.ServeJSON()
	}

	//生成 token，并存放到缓存中
	token := uuid.Must(uuid.NewV1()).String()
	err = userCache.Put(token, u)
	if err != nil {
		login.Data["json"] = e.Generate(e.AuthError, err)
		login.ServeJSON()
	}
	login.Data["json"] = e.Generate(e.Success, map[string]string{"token": token})
	login.ServeJSON()
}

