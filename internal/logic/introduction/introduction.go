package introduction

import (
	"context"
	"database/sql"
	"gf_blog/internal/dao"
	"gf_blog/internal/model"
	"gf_blog/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
)

func init() {
	service.RegisterIntroduction(New())
}

func New() *sIntroduction {
	return &sIntroduction{}
}

type sIntroduction struct{}

func (s *sIntroduction) GetByUserId(ctx context.Context, in *model.GetIntroductionByUserIdInput) (out *model.GetIntroductionByUserIdOutput, err error) {
	out = &model.GetIntroductionByUserIdOutput{}
	if err := dao.BlogProfile.Ctx(ctx).Where(dao.BlogProfile.Columns().UserId, in.UserId).Scan(out); err != nil {
		if err == sql.ErrNoRows {
			return nil, gerror.New("没有自我介绍内容")
		}
		return nil, err
	}
	return
}
