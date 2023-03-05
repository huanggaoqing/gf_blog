package file

import (
	"context"
	"gf_blog/internal/model"
	"gf_blog/internal/service"
	"gf_blog/utility"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

type sFile struct{}

func (s *sFile) Upload(ctx context.Context, in *model.FileUploadInput) (out *model.FileUploadOutput, err error) {
	// 获取图片上传路径
	var pathKey string
	switch in.Type {
	case "project":
		pathKey = "upload.projectImagePath"
	case "avatar":
		pathKey = "upload.avatarImagePath"
	}
	uploadPath := g.Cfg().MustGet(ctx, pathKey).String()
	if uploadPath == "" {
		return nil, gerror.New("获取上传路径失败")
	}
	path := gfile.Join(uploadPath)
	fileName, err := in.File.Save(path, true)
	if err != nil {
		return nil, err
	}
	url, err := utility.GetLocalIp(ctx)
	if err != nil {
		return nil, err
	}
	g.Log().Debugf(ctx, "当前ip：%s", url)
	return &model.FileUploadOutput{Url: url + "/" + gfile.Join(path, fileName)}, nil
}
