package consts

const (
	SUCCESS = 0

	ERROR = 10000
	// JWT错误
	AUTH_ERROR     = 10001
	JWT_AUTH_ERROR = 10002
)

var ERROR_MAP = map[int]string{
	SUCCESS: "ok",

	AUTH_ERROR:     "验证错误",
	JWT_AUTH_ERROR: "JWT验证错误",
}

func GetMsg(code int) string {
	if msg, ok := ERROR_MAP[code]; ok {
		return msg
	}
	return ERROR_MAP[ERROR]
}
