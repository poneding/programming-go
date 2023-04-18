package e

const (
	Success = 200

	Error         = 500
	InvalidParams = 400

	ErrorExistUser    = 10001
	ErrorNotExistUser = 10003
	ErrorFaileEncrypt = 10006
	ErrorNotCompare   = 10007

	HaveSignup           = 20001
	ErrorActivityTimeout = 20002

	ErrorAuthCheckTokenFail    = 30001 //token 错误
	ErrorAuthCheckTokenTimeout = 30002 //token 过期
	ErrorAuthToken             = 30003
	ErrorAuth                  = 30004
	ErrorAuthNotFound          = 30005
	ErrorDatabase              = 40001
)
