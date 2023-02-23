// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BlogTag is the golang structure of table blog_tag for DAO operations like Where/Data.
type BlogTag struct {
	g.Meta     `orm:"table:blog_tag, do:true"`
	Id         interface{} //
	Name       interface{} //
	Url        interface{} //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
}