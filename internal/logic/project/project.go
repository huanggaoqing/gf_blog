package project

import (
	"context"
	"gf_blog/internal/dao"
	"gf_blog/internal/model"
	"gf_blog/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterProject(New())
}

func New() *sProject {
	return &sProject{}
}

type sProject struct{}

func (s *sProject) GetListById(ctx context.Context, in *model.GetProjectListByIdInput) (out *model.GetProjectListByUserIdOutput, err error) {
	out = &model.GetProjectListByUserIdOutput{
		Page: in.Page,
		Size: in.Size,
	}
	m := dao.BlogProject.Ctx(ctx)
	list := m.Where(dao.BlogProject.Columns().UserId, in.UserId)
	out.Total, err = list.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if err = list.Page(in.Page, in.Size).Scan(&out.List); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return
}
