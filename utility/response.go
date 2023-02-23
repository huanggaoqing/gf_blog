package utility

import "github.com/gogf/gf/v2/net/ghttp"

type response[T interface{}] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func JsonExit[T interface{}](r *ghttp.Request, code int, data T) {
	r.Response.WriteJsonExit(response[T]{
		Code: code,
		Data: data,
		Msg:  "success",
	})
}

func FailExit(r *ghttp.Request, code int, msg string) {
	r.Response.WriteJsonExit(response[string]{
		Code: code,
		Data: "",
		Msg:  msg,
	})
}
