package consts

const (
	// for user
	CtxUserId     = "CtxUserId"
	CtxUserName   = "CtxUserName"
	CtxUserAvatar = "CtxUserAvatar"
	//for admin
	CtxAdminId       = "CtxAdminId"
	CtxAdminName     = "CtxAdminName"
	CtxAdminRoleCode = "CtxAdminRoleCode"
	// for login
	FrontendKey        = "frontend:"
	BackendKey         = "backend:"
	TokenType          = "Bearer"
	CacheModeRedis     = 2
	BackendServerName  = "个人网站server"
	FrontendMultiLogin = false
	GTokenExpireIn     = 10 * 24 * 60 * 60
)
