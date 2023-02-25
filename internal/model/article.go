package model

import (
	"gf_blog/internal/model/entity"
)

type GetArticleListItemOutput struct {
	entity.BlogArticle
	UserName string
}

type GetArticleListInput struct {
	UserId string
	Page   int
	Size   int
}

type GetArticleListOutput struct {
	Page  int
	Size  int
	List  []*GetArticleListItemOutput
	Total int
}

type GetArticleDetailInput struct {
	ArticleId string
}

type GetArticleDetailOutput struct {
	entity.BlogArticle
	UserName string
	Tag      []*entity.BlogTag
}

type GetTagDataByArticleIdInput struct {
	ArticleId string
}

type GetTagDataByArticleIdOutput struct {
	List []*entity.BlogTag
}
