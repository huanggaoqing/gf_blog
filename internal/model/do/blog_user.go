// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BlogUser is the golang structure of table blog_user for DAO operations like Where/Data.
type BlogUser struct {
	g.Meta      `orm:"table:blog_user, do:true"`
	UserId      interface{} //
	UserName    interface{} //
	Password    interface{} //
	PhoneNumber interface{} //
	UserType    interface{} //
	OpenId      interface{} //
	Avatar      interface{} //
	Role        interface{} //
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
}
