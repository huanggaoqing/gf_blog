package main

import (
	"gf_blog/internal/cmd"
	_ "gf_blog/internal/logic"
	_ "gf_blog/internal/packed"
	"gf_blog/utility"
	_ "gf_blog/utility"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"log"
)

func main() {
	log.Println(utility.IdBySnowflake.Generate())
	cmd.Main.Run(gctx.New())
}
