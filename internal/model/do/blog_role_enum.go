// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BlogRoleEnum is the golang structure of table blog_role_enum for DAO operations like Where/Data.
type BlogRoleEnum struct {
	g.Meta     `orm:"table:blog_role_enum, do:true"`
	RoleId     interface{} //
	RoleName   interface{} //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
}