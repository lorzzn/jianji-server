package r

const (
	OK   = 0
	FAIL = 500

	// code= 9000... 通用错误
	ERROR_BAD_PARAM           = 9001
	ERROR_REQUEST_PAGE        = 9002
	ERROR_INVALID_PARAM       = 9003
	ERROR_DB_OPE              = 9004
	ERROR_DB_DATA_NOT_EXISTED = 9005

	// token 错误
	TOKEN_IS_EMPTY              = 2003
	TOKEN_IS_BAD                = 2004
	TOKEN_AUTHORIZATION_INVALID = 2005
	TOKEN_NOT_EXISTED           = 2006
	TOKEN_AUTHORIZATION_FAILED  = 2007

	//应用接口错误
	APP_CREATERSA_FAILED = 20001
	APP_SAVERSA_FAILED   = 20002
	APP_GETRSA_FAILED    = 20003

	//用户接口错误
	USER_EXISTED             = 10001
	USER_PASSWORD_INCORRECT  = 10002
	USER_NOT_LOGIN           = 10003
	USER_REFRESHTOKEN_FAILED = 10004
	USER_NOT_EXISTED         = 10005
	USER_OPE_NOT_ALLOW       = 10006
)

var codeMsg = map[int]string{
	OK:   "OK",
	FAIL: "FAIL",

	ERROR_BAD_PARAM:           "请求参数格式错误",
	ERROR_REQUEST_PAGE:        "分页参数错误",
	ERROR_INVALID_PARAM:       "不合法的请求参数",
	ERROR_DB_OPE:              "数据库操作异常",
	ERROR_DB_DATA_NOT_EXISTED: "数据不存在",

	TOKEN_IS_EMPTY:              "TOKEN 不能为空",
	TOKEN_IS_BAD:                "TOKEN 格式错误",
	TOKEN_AUTHORIZATION_INVALID: "登录已过期",
	TOKEN_NOT_EXISTED:           "未登录",
	TOKEN_AUTHORIZATION_FAILED:  "登录失败",

	//应用接口错误
	APP_CREATERSA_FAILED: "生成授权密钥出错",
	APP_SAVERSA_FAILED:   "保存授权密钥出错",
	APP_GETRSA_FAILED:    "获取授权密钥出错",

	//用户接口错误
	USER_EXISTED:             "用户已注册，请直接登录",
	USER_PASSWORD_INCORRECT:  "登录密码错误",
	USER_NOT_LOGIN:           "用户未登录",
	USER_REFRESHTOKEN_FAILED: "登录已过期，请重新登录",
	USER_NOT_EXISTED:         "用户不存在",
	USER_OPE_NOT_ALLOW:       "操作不允许",
}

func GetCodeMsg(code int) string {
	return codeMsg[code]
}
