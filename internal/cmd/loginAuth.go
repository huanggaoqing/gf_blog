package cmd

import (
	"context"
	"gf_blog/api/frontend"
	"gf_blog/internal/consts"
	"gf_blog/internal/dao"
	"gf_blog/internal/model/entity"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type response[T interface{}] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func StartFrontendToken() (frontendToken *gtoken.GfToken, err error) {
	frontendToken = &gtoken.GfToken{
		ServerName:      consts.BackendServerName,
		LoginPath:       "/user/login",
		LoginBeforeFunc: frontendLoginBefore,
		LoginAfterFunc:  frontendLoginAfter,
		LogoutPath:      "/user/logout",
		AuthAfterFunc:   frontendAuthAfter,
		CacheMode:       consts.CacheModeRedis,
		MultiLogin:      consts.FrontendMultiLogin,
	}
	return
}

// 登录验证方法 return userKey 用户标识 如果userKey为空，结束执行
func frontendLoginBefore(r *ghttp.Request) (string, interface{}) {
	name := r.Get("userName").String()
	password := r.Get("password").String()
	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("登录失败，账号或密码错误"))
		r.ExitAll()
	}
	user := &entity.BlogUser{}
	err := dao.BlogUser.Ctx(context.TODO()).
		Where(dao.BlogUser.Columns().UserName, name).
		Where(dao.BlogUser.Columns().Password, password).
		Scan(user)
	if err != nil {
		return "", err
	}
	return consts.FrontendKey + user.UserId, user
}

// 登录返回方法
func frontendLoginAfter(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Success() {
		userKey := gstr.StrEx(respData.Get("userKey").String(), consts.FrontendKey)
		// 通过userId获取用户信息
		user := &entity.BlogUser{}
		dao.BlogUser.Ctx(context.TODO()).Where(dao.BlogUser.Columns().UserId, userKey).Scan(user)
		data := &frontend.UserLoginRes{}
		if err := gconv.Struct(user, data); err != nil {
			return
		}
		data.TokenType = consts.TokenType
		data.Token = respData.Get("token").String()
		data.ExpireIn = consts.GTokenExpireIn
		r.Response.WriteJson(&response[*frontend.UserLoginRes]{
			Code: 0,
			Data: data,
			Msg:  "success",
		})
		r.Exit()
	} else {
		respData.Code = 0
		r.Response.WriteJson(respData)
	}
}

// 认证返回方法
func frontendAuthAfter(r *ghttp.Request, respData gtoken.Resp) {
	user := &entity.BlogUser{}
	if err := gconv.Struct(respData.GetString("data"), user); err != nil {
		r.Response.WriteJson(&response[string]{
			Code: 403,
			Data: "",
			Msg:  "token无效或已过期",
		})
		return
	}
	//todo 这里可以写账号前置校验、是否被拉黑、有无权限等逻辑
	r.SetCtxVar(consts.CtxUserId, user.UserId)
	r.SetCtxVar(consts.CtxUserName, user.UserName)
	r.SetCtxVar(consts.CtxUserAvatar, user.Avatar)
	r.Middleware.Next()
}
