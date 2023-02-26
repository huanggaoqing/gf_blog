// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BlogProject is the golang structure of table blog_project for DAO operations like Where/Data.
type BlogProject struct {
	g.Meta          `orm:"table:blog_project, do:true"`
	ProjectId       interface{} // 项目Id
	ProjectName     interface{} // 项目名称
	ProjectUrl      interface{} // 项目地址
	ProjectDesc     interface{} // 项目描述
	ProjectIcon     interface{} // 项目图标
	ProjectIconDark interface{} // 项目图标(暗色模式)
	UserId          interface{} // 用户id
	CreateTime      *gtime.Time // 创建时间
	UpdateTime      *gtime.Time // 更新时间
}