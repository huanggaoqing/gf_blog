// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BlogTagAssociate is the golang structure for table blog_tag_associate.
type BlogTagAssociate struct {
	Id         int         `json:"id"         ` //
	ArticleId  string      `json:"articleId"  ` //
	TagId      int         `json:"tagId"      ` //
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
}
