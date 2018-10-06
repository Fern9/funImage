package e

type e struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	//成功
	Success = "success"

	//一般性错误定义

	//参数错误
	ParamsError="paramsError"

	//用户未登录
	NotAuth = "notAuth"
	//用户登录失败
	AuthError = "authError"
	GetUserInfoError = "getUserInfoError"

	//获取微信小程序用户信息失败
	GetMiniProgramUserInfoError = "getMiniProgramUserInfoError"
)

var elist = map[string]e{
	"success":   {0, "success", nil},

	"paramsError": {1001, "参数错误", nil},

	"notAuth":   {2001, " 用户未登录", nil},
	"authError": {2002, " 用户登录失败", nil},
	"getUserInfoError": {2003, "获取用户信息失败", nil},

	"getMiniProgramUserInfoError": {3001, "获取微信小程序用户信息失败", nil},
}
