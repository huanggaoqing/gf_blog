package frontend

import "github.com/gogf/gf/v2/frame/g"

type UserInfoBase struct {
	UserId      string `json:"userId"`
	UserName    string `json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
	UserType    string `json:"userType"`
	OpenId      string `json:"openId"`
	Avatar      string `json:"avatar"`
	Role        string `json:"role"`
}

// RegisterReq 用户注册
type RegisterReq struct {
	g.Meta      `path:"/user/register" tags:"用户" method:"post" summary:"用户注册"`
	PhoneNumber string `json:"phoneNumber" v:"required#手机号不可为空, phone#手机号不符合规范" dc:"手机号"`
	UserName    string `json:"userName" v:"required#用户名不可为空" dc:"用户名"`
	Password    string `json:"password" v:"required#密码不可为空, password2#密码不符合规范" dc:"密码"`
	Password2   string `json:"password2" v:"required#密码不可为空, password2#密码不符合规范" dc:"确认密码"`
	Avatar      string `json:"avatar" v:"required#头像不可为空" dc:"头像"`
}

type RegisterRes struct {
	UserId string `json:"userId"`
}

// UserLoginRes 用户登录
type UserLoginRes struct {
	TokenType string `json:"tokenType"`
	Token     string `json:"token"`
	ExpireIn  int64  `json:"expireIn"`
	UserInfoBase
}
