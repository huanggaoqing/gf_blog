package controller

import (
	"context"
	"gf_blog/api/frontend"
	"gf_blog/internal/model"
	"gf_blog/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var Project = cProject{}

type cProject struct{}

func (c *cProject) GetListByUserId(ctx context.Context, req *frontend.GetProjectListByUserIdReq) (res *frontend.GetProjectListByUserIdRes, err error) {
	params := &model.GetProjectListByIdInput{}
	if err = gconv.Struct(req, params); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	data, err := service.Project().GetListById(ctx, params)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	res = &frontend.GetProjectListByUserIdRes{}
	if err = gconv.Struct(data, res); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return
}
