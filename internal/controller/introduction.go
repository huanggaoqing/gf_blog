package controller

import (
	"context"
	"gf_blog/api/frontend"
	"gf_blog/internal/model"
	"gf_blog/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

var Introduction = cIntroduction{}

type cIntroduction struct{}

func (c *cIntroduction) GetByUserId(ctx context.Context, req *frontend.GetIntroductionInfoReq) (res *frontend.GetIntroductionInfoRes, err error) {
	data := &model.GetIntroductionByUserIdInput{}
	if err = gconv.Struct(req, data); err != nil {
		return nil, err
	}
	introduction, err := service.Introduction().GetByUserId(ctx, data)
	if err != nil {
		return nil, err
	}
	res = &frontend.GetIntroductionInfoRes{}
	if err = gconv.Struct(introduction, res); err != nil {
		return nil, err
	}
	return
}
