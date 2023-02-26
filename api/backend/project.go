package backend

import "github.com/gogf/gf/v2/frame/g"

type AddProjectReq struct {
	g.Meta      `path:"/project/add" tags:"项目" method:"post" summary:"新增用户项目"`
	ProjectId   string `json:"projectId" v:"required#projectId不可为空" dc:"项目Id"'`
	ProjectName string `json:"projectName" v:"required#projectName不可为空" dc:"项目名称"`
	ProjectUrl  string `json:"projectUrl" v:"required#projectUrl不可为空" dc:"项目地址"`
	ProjectDesc string `json:"projectDesc" v:"required#projectDesc不可为空" dc:"项目描述"`
	ProjectIcon string `json:"projectIcon" v:"required#projectIcon不可为空" dc:"项目图标"`
}

type AddProjectRes struct {
	ProjectId string `json:"projectId"`
}
