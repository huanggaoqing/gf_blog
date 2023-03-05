package cmd

import (
	"context"
	"gf_blog/internal/cmd/loginAuth"
	"gf_blog/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf_blog/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			test := g.Cfg().MustGet(ctx, "test").String()
			g.Log().Debug(ctx, "当前环境", test)
			s := g.Server()
			s.AddStaticPath("/upload", "/upload")
			frontendToken, err := loginAuth.StartFrontendToken()
			if err != nil {
				return err
			}
			// 客户端
			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					ghttp.MiddlewareHandlerResponse,
				)
				// 不需要token鉴权的路由
				group.Bind(
					controller.User.Register,
					controller.Introduction.GetByUserId,
					controller.Article.GetList,
					controller.Article.GetDetail,
					controller.Project.GetListByUserId,
					controller.File.Upload,
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					// 启动token鉴权
					err := frontendToken.Middleware(ctx, group)
					if err != nil {
						return
					}
					// 需要token鉴权的路由
					group.Bind()
				})
			})
			backendToken, err := loginAuth.StartBackendToken()
			if err != nil {
				return err
			}
			// 后台管理
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					ghttp.MiddlewareHandlerResponse,
				)
				// 不需要token鉴权的路由
				group.Bind()
				group.Group("/", func(group *ghttp.RouterGroup) {
					// 启动token鉴权
					err := backendToken.Middleware(ctx, group)
					if err != nil {
						return
					}
					group.Bind()
				})
			})
			s.Run()
			return nil
		},
	}
)
