package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ProjectItemBase struct {
	ProjectId       string      `json:"projectId"   `     // 项目Id
	ProjectName     string      `json:"projectName" `     // 项目名称
	ProjectUrl      string      `json:"projectUrl"  `     // 项目地址
	ProjectDesc     string      `json:"projectDesc" `     // 项目描述
	ProjectIcon     string      `json:"projectIcon" `     // 项目图标
	ProjectIconDark string      `json:"projectIconDark" ` // 项目图标
	UserId          string      `json:"userId"      `     // 用户id
	CreateTime      *gtime.Time `json:"createTime"  `     // 创建时间
	UpdateTime      *gtime.Time `json:"updateTime"  `     // 更新时间
}

type GetProjectListByUserIdReq struct {
	g.Meta `path:"/project/getByUserId" tags:"项目" method:"get" summary:"获取用户项目列表"`
	UserId string `json:"userId" v:"required#userId不可为空" dc:"用户ID"`
	Page   int    `json:"page" v:"required#page不可为空" dc:"页码"`
	Size   int    `json:"size" v:"required#size不可为空" dc:"每页数量"`
}

type GetProjectListByUserIdRes struct {
	Page  int                `json:"page"`
	Size  int                `json:"size"`
	List  []*ProjectItemBase `json:"list"`
	Total int                `json:"total"`
}
