package controller

import (
	"context"
	"gf_blog/api/frontend"
	"gf_blog/internal/model"
	"gf_blog/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

var User = cUser{}

type cUser struct{}

func (c *cUser) Register(ctx context.Context, req *frontend.RegisterReq) (res *frontend.RegisterRes, err error) {
	data := &model.RegisterInput{}
	err = gconv.Struct(req, data)
	if err != nil {
		return nil, err
	}
	user, err := service.User().Register(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.RegisterRes{UserId: user.UserId}, nil
}

//func (c *cUser) Login(ctx context.Context, req *frontend.UserLoginReq) (res *frontend.UserLoginRes, err error) {
//	data := &model.LoginInput{}
//	err = gconv.Struct(req, data)
//	if err != nil {
//		return nil, err
//	}
//	_, err = service.User().Login(ctx, data)
//	if err != nil {
//		return nil, err
//	}
//	return
//}
