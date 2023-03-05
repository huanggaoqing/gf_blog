package backend

type UserInfoBase struct {
	UserId      string `json:"userId"`
	UserName    string `json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
	UserType    string `json:"userType"`
	OpenId      string `json:"openId"`
	Avatar      string `json:"avatar"`
}

type AdminInfo struct {
	UserInfoBase
	RoleCode int `json:"roleCode"`
}

// UserLoginRes 用户登录
type UserLoginRes struct {
	TokenType string `json:"tokenType"`
	Token     string `json:"token"`
	ExpireIn  int64  `json:"expireIn"`
	AdminInfo
}
