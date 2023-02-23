// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BlogArticle is the golang structure for table blog_article.
type BlogArticle struct {
	ArticleId      int64       `json:"articleId"      ` //
	UserId         int64       `json:"userId"         ` //
	ArticleContent string      `json:"articleContent" ` //
	ArticleTitle   string      `json:"articleTitle"   ` //
	IsStick        int64       `json:"isStick"        ` //
	GroupId        int64       `json:"groupId"        ` //
	CreateTime     *gtime.Time `json:"createTime"     ` //
	UpdateTime     *gtime.Time `json:"updateTime"     ` //
}
