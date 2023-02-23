// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BlogUser is the golang structure for table blog_user.
type BlogUser struct {
	UserId      string      `json:"userId"      ` //
	UserName    string      `json:"userName"    ` //
	Password    string      `json:"password"    ` //
	PhoneNumber string      `json:"phoneNumber" ` //
	UserType    string      `json:"userType"    ` //
	OpenId      string      `json:"openId"      ` //
	Avatar      string      `json:"avatar"      ` //
	Role        int64       `json:"role"        ` //
	CreateTime  *gtime.Time `json:"createTime"  ` //
	UpdateTime  *gtime.Time `json:"updateTime"  ` //
}