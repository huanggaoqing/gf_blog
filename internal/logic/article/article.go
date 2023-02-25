package article

import (
	"context"
	"database/sql"
	"fmt"
	"gf_blog/internal/dao"
	"gf_blog/internal/model"
	"gf_blog/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterArticle(New())
}

func New() *sArticle {
	return &sArticle{}
}

type sArticle struct{}

func (s *sArticle) GetList(ctx context.Context, in *model.GetArticleListInput) (out *model.GetArticleListOutput, err error) {
	out = &model.GetArticleListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	var (
		articleTable   = dao.BlogArticle.Table()
		articleColumns = dao.BlogArticle.Columns()
		userTable      = dao.BlogUser.Table()
		userColumns    = dao.BlogUser.Columns()
	)
	m := dao.BlogArticle.Ctx(ctx)
	list := m.FieldsPrefix(articleTable, articleColumns).
		FieldsPrefix(userTable, userColumns.UserName).
		LeftJoinOnField(userTable, userColumns.UserId).
		WherePrefix(articleTable, articleColumns.UserId, in.UserId).
		OrderDesc(articleColumns.IsStick).
		Order(fmt.Sprintf("%v.%v %v", articleTable, articleColumns.CreateTime, "DESC"))
	out.Total, err = m.Where(articleColumns.UserId, in.UserId).Count()
	if err != nil {
		return nil, err
	}
	if err := list.Page(in.Page, in.Size).Scan(&out.List); err != nil {
		if err == sql.ErrNoRows {
			return nil, gerror.New("文章列表为空")
		}
		return nil, err
	}
	return
}

func (s *sArticle) GetDetail(ctx context.Context, in *model.GetArticleDetailInput) (out *model.GetArticleDetailOutput, err error) {
	out = &model.GetArticleDetailOutput{}
	var (
		articleTable   = dao.BlogArticle.Table()
		articleColumns = dao.BlogArticle.Columns()
		userTable      = dao.BlogUser.Table()
		userColumns    = dao.BlogUser.Columns()
	)
	m := dao.BlogArticle.Ctx(ctx)
	// TODO 记录blog
	list := m.FieldsPrefix(articleTable, articleColumns).
		FieldsPrefix(userTable, userColumns.UserName).
		LeftJoinOnField(userTable, userColumns.UserId).
		WherePrefix(articleTable, articleColumns.ArticleId, in.ArticleId)
	if err = list.Scan(out); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	tagData, err := s.GetTagDataByArticleId(ctx, &model.GetTagDataByArticleIdInput{
		ArticleId: in.ArticleId,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	out.Tag = tagData.List
	return
}

func (s *sArticle) GetTagDataByArticleId(ctx context.Context, in *model.GetTagDataByArticleIdInput) (out *model.GetTagDataByArticleIdOutput, err error) {
	articleIds := make([]int, 0)
	// TODO 记录blog
	tagColumns := dao.BlogTagAssociate.Columns()
	res, err := dao.BlogTagAssociate.Ctx(ctx).
		Fields(tagColumns.TagId).
		Where(tagColumns.ArticleId, in.ArticleId).
		All()
	for _, v := range res {
		value := v[tagColumns.TagId]
		articleIds = append(articleIds, value.Int())
	}
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	out = &model.GetTagDataByArticleIdOutput{}
	if err = dao.BlogTag.Ctx(ctx).WhereIn(dao.BlogTag.Columns().Id, articleIds).Scan(&out.List); err != nil {
		return nil, err
	}
	return
}
