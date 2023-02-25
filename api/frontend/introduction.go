package frontend

import "github.com/gogf/gf/v2/frame/g"

// 获取用户自我介绍
type GetIntroductionInfoReq struct {
	g.Meta `path:"/introduction/getByUserId" tags:"自我介绍" method:"get" summary:"获取用户自我介绍"`
	UserId string `json:"userId" v:"required#userId不可为空" dc:"用户ID"`
}

type GetIntroductionInfoRes struct {
	IntroductionId string `json:"introductionId"`
	UserId         string `json:"userId"`
	Content        string `json:"content"`
	CreateTime     string `json:"createTime"`
	UpdateTime     string `json:"updateTime"`
}
