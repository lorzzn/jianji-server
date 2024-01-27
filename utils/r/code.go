package r

const (
	OK   = 0
	FAIL = 500

	// code= 9000... 通用错误
	ERROR_BAD_PARAM     = 9001
	ERROR_REQUEST_PAGE  = 9002
	ERROR_INVALID_PARAM = 9003
	ERROR_DB_OPE        = 9004

	//用户接口错误
	USER_EXISTED = 10001
)

var codeMsg = map[int]string{
	OK:   "OK",
	FAIL: "FAIl",

	ERROR_BAD_PARAM:     "请求参数格式错误",
	ERROR_REQUEST_PAGE:  "分页参数错误",
	ERROR_INVALID_PARAM: "不合法的请求参数",
	ERROR_DB_OPE:        "数据库操作异常",

	//用户接口错误
	USER_EXISTED: "用户已注册，请直接登录",
}

func GetCodeMsg(code int) string {
	return codeMsg[code]
}
