package utility

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"log"
	"net"
)

func GetLocalIp(ctx context.Context) (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	port := g.Cfg().MustGet(ctx, "server.address").String()
	if port == "" {
		return "", gerror.New("获取端口失败")
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	address := fmt.Sprintf("http://%s%s", localAddr.IP.String(), port)
	return address, nil
}
