package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/file/upload" tags:"文件" method:"post" summary:"文件上传"`
	File   *ghttp.UploadFile `json:"file" v:"required#file不可为空" dc:"上传的文件"`
	Type   string            `json:"type" v:"required#type不可为空" dc:"上传文件的类型"`
}

type FileUploadRes struct {
	Url string `json:"url"`
}
