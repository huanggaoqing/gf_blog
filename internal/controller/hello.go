package controller

import (
	"context"
	"log"

	"github.com/gogf/gf/v2/frame/g"

	"gf_blog/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	value, err := g.Redis().Get(ctx, "name")
	if err != nil {
		return nil, err
	}
	log.Println(value.String())
	set, err := g.Redis().Set(ctx, "age", 21)
	if err != nil {
		return nil, err
	}
	log.Println(set.String())
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
