package controller

import (
	"context"
	"gf_blog/api/frontend"
	"gf_blog/internal/model"
	"gf_blog/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

var Article = cArticle{}

type cArticle struct{}

func (c *cArticle) GetList(ctx context.Context, req *frontend.GetArticleListReq) (res *frontend.GetArticleListRes, err error) {
	params := &model.GetArticleListInput{}
	if err := gconv.Struct(req, params); err != nil {
		return nil, err
	}
	data, err := service.Article().GetList(ctx, params)
	if err != nil {
		return nil, err
	}
	res = &frontend.GetArticleListRes{}
	if err = gconv.Struct(data, res); err != nil {
		return nil, err
	}
	return
}

func (c *cArticle) GetDetail(ctx context.Context, req *frontend.GetArticleDetailByIdReq) (res *frontend.GetArticleDetailByIdRes, err error) {
	detail, err := service.Article().GetDetail(ctx, &model.GetArticleDetailInput{ArticleId: req.ArticleId})
	if err != nil {
		return nil, err
	}
	res = &frontend.GetArticleDetailByIdRes{}
	if err = gconv.Struct(detail, res); err != nil {
		return nil, err
	}
	return
}
