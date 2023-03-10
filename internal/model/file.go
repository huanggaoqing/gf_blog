package model

import "github.com/gogf/gf/v2/net/ghttp"

type FileUploadInput struct {
	File *ghttp.UploadFile //上传文件对象
	Type string
}

type FileUploadOutput struct {
	Url string
}
