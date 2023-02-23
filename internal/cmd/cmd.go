package cmd

import (
	"context"
	"gf_blog/internal/cmd/loginAuth"

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
			s := g.Server()
			frontendToken, err := loginAuth.StartFrontendToken()
			if err != nil {
				return err
			}
			// 客户端
			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				// 不需要token鉴权的路由
				group.Bind(
					controller.User.Register,
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
			s.Run()
			return nil
		},
	}
)
