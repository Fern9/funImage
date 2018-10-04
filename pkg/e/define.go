package e

type e struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

const(
	//成功
	SUCCESS="SUCCESS"

	//用户未登录
	NOT_AUTH="NOT_AUTH"
)

var elist = map[string]e{
	"SUCCESS": {0, "success", nil},
	"NOT_AUTH": {2001, " 用户未登录", nil},
}



