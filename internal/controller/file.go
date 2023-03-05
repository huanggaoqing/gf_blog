package controller

import (
	"context"
	"gf_blog/api/frontend"
	"gf_blog/internal/model"
	"gf_blog/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

var File = cFile{}

type cFile struct{}

func (c *cFile) Upload(ctx context.Context, req *frontend.FileUploadReq) (res *frontend.FileUploadRes, err error) {
	params := &model.FileUploadInput{}
	if err = gconv.Struct(req, params); err != nil {
		return nil, err
	}
	data, err := service.File().Upload(ctx, params)
	if err != nil {
		return nil, err
	}
	return &frontend.FileUploadRes{Url: data.Url}, err
}
