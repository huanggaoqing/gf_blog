package loginAuth

import (
	"context"
	"gf_blog/api/backend"
	"gf_blog/internal/consts"
	"gf_blog/internal/dao"
	"gf_blog/utility"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func StartBackendToken() (backendToken *gtoken.GfToken, err error) {
	backendToken = &gtoken.GfToken{
		ServerName:      consts.BackendServerName,
		LoginPath:       "/user/login",
		LoginBeforeFunc: backendLoginBefore,
		LoginAfterFunc:  backendLoginAfter,
		LogoutPath:      "/user/logout",
		AuthAfterFunc:   backendAuthAfter,
		CacheMode:       consts.CacheModeRedis,
		MultiLogin:      consts.FrontendMultiLogin,
	}
	return
}

// 登录验证方法 return userKey 用户标识 如果userKey为空，结束执行
func backendLoginBefore(r *ghttp.Request) (string, interface{}) {
	name := r.Get("userName").String()
	password := r.Get("password").String()
	if name == "" || password == "" {
		utility.FailExit(r, -1, "登录失败，账号或密码错误")
	}
	var (
		userTable   = dao.BlogUser.Table()
		userColumns = dao.BlogUser.Columns()
		roleTable   = dao.BlogRole.Table()
		roleColumns = dao.BlogRole.Columns()
	)
	info := &backend.AdminInfo{}
	m := dao.BlogUser.Ctx(context.TODO())
	err := m.FieldsPrefix(userTable, userColumns).
		FieldsPrefix(roleTable, roleColumns.RoleCode).
		LeftJoinOnField(roleTable, roleColumns.UserId).
		WherePrefix(userTable, userColumns.UserName, name).
		WherePrefix(userTable, userColumns.Password, password).
		WherePrefix(roleTable, roleColumns.RoleCode, 1).
		Scan(info)
	if err != nil {
		return "", err
	}
	if info.RoleCode != 1 {
		utility.FailExit(r, -1, "登录失败，此账号无权限")
	}
	return consts.BackendKey + info.UserId, info
}

// 登录返回方法
func backendLoginAfter(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Success() {
		userKey := gstr.StrEx(respData.Get("userKey").String(), consts.BackendKey)
		// 通过userId获取用户信息
		user := &backend.AdminInfo{}
		var (
			userTable   = dao.BlogUser.Table()
			userColumns = dao.BlogUser.Columns()
			roleTable   = dao.BlogRole.Table()
			roleColumns = dao.BlogRole.Columns()
		)
		m := dao.BlogUser.Ctx(context.TODO())
		err := m.FieldsPrefix(userTable, userColumns).
			FieldsPrefix(roleTable, roleColumns.RoleCode).
			LeftJoinOnField(roleTable, roleColumns.UserId).
			WherePrefix(userTable, userColumns.UserId, userKey).
			Scan(user)
		if err != nil {
			utility.FailExit(r, 1, "登录失败")
		}
		data := &backend.UserLoginRes{}
		if err := gconv.Struct(user, data); err != nil {
			return
		}
		data.TokenType = consts.TokenType
		data.Token = respData.Get("token").String()
		data.ExpireIn = consts.GTokenExpireIn
		utility.JsonExit[*backend.UserLoginRes](r, 0, data)
	} else {
		respData.Code = 0
		r.Response.WriteJson(respData)
	}
}

// 认证返回方法
func backendAuthAfter(r *ghttp.Request, respData gtoken.Resp) {
	user := &backend.AdminInfo{}
	if err := gconv.Struct(respData.GetString("data"), user); err != nil {
		utility.FailExit(r, 403, "token无效或已过期")
	}
	//todo 这里可以写账号前置校验、是否被拉黑、有无权限等逻辑
	r.SetCtxVar(consts.CtxAdminId, user.UserId)
	r.SetCtxVar(consts.CtxAdminName, user.UserName)
	r.SetCtxVar(consts.CtxAdminRoleCode, user.RoleCode)
	r.Middleware.Next()
}
