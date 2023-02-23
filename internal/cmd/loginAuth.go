package cmd

import (
	"context"
	"gf_blog/internal/consts"
	"gf_blog/internal/dao"
	"gf_blog/internal/model/entity"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"log"
)

func StartFrontendToken() (frontendToken *gtoken.GfToken, err error) {
	frontendToken = &gtoken.GfToken{
		LoginPath:       "/user/login",
		LoginBeforeFunc: frontendLoginBefore,
		LoginAfterFunc:  frontendLoginAfter,
		LogoutPath:      "/user/logout",
	}
	return
}

func frontendLoginBefore(r *ghttp.Request) (string, interface{}) {
	name := r.Get("userName").String()
	password := r.Get("password").String()
	if name == "" || password == "" {
		r.Response.WriteJson("登录失败，账号或密码错误")
		r.ExitAll()
	}
	user := &entity.BlogUser{}
	err := dao.BlogUser.Ctx(context.TODO()).
		Where(dao.BlogUser.Columns().UserName, name).
		Where(dao.BlogUser.Columns().Password, password).
		Scan(user)
	log.Printf("%+v", user)
	if err != nil {
		return "", err
	}
	return consts.FRONTENDKEY + user.UserId, user
}

func frontendLoginAfter(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Success() {
		userKey := gstr.StrEx(respData.Get("userKey").String(), consts.FRONTENDKEY)
		log.Println("userKey", userKey)
	} else {
		respData.Code = 0
		r.Response.WriteJson(respData)
	}
}
