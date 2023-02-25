package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ArticleBase struct {
	ArticleId    string      `json:"articleId"`
	UserId       string      `json:"userId"`
	UserName     string      `json:"userName"`
	ArticleTitle string      `json:"articleTitle"`
	IsStick      int64       `json:"isStick"`
	CreateTime   *gtime.Time `json:"createTime"`
	UpdateTime   *gtime.Time `json:"updateTime"`
}

type GetArticleListReq struct {
	g.Meta `path:"/article/getList" tags:"文章" method:"get" summary:"获取用户文章列表"`
	UserId string `json:"userId" v:"required#userId不可为空" dc:"用户ID"`
	Page   int    `json:"page" v:"required#page不可为空" dc:"页码"`
	Size   int    `json:"size" v:"required#size不可为空" dc:"每页数量"`
}

type GetArticleListRes struct {
	Page  int            `json:"page"`
	Size  int            `json:"size"`
	List  []*ArticleBase `json:"list"`
	Total int            `json:"total"`
}

type GetArticleDetailByIdReq struct {
	g.Meta    `path:"/article/detail" tags:"文章" method:"get" summary:"获取用户文章详情"`
	ArticleId string `json:"articleId" v:"required#articleId不可为空" dc:"文章Id"`
}

type GetArticleDetailByIdRes struct {
	ArticleContent string            `json:"articleContent" ` //
	GroupId        int64             `json:"groupId"        ` //
	Tag            []*ArticleTagBase `json:"tag"`
	ArticleBase
}
