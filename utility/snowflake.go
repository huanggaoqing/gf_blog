package utility

import (
	"fmt"
	"github.com/housemecn/snowflake"
	"log"
)

var IdBySnowflake *generateId

func init() {
	newSnowflake, err := snowflake.NewSnowflake(int64(0), int64(0))
	if err != nil {
		log.Panicln("雪花算法初始化失败")
		return
	}
	IdBySnowflake = &generateId{s: newSnowflake}
}

type generateId struct {
	s *snowflake.Snowflake
}

func (g *generateId) Generate() string {
	id := g.s.NextVal()
	return fmt.Sprintf("%v", id)
}
