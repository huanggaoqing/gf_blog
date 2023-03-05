package utility

import "github.com/gogf/gf/v2/net/ghttp"

type response[T interface{}] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func JsonExit[T interface{}](r *ghttp.Request, code int, data T) {
	r.Response.WriteJsonExit(response[T]{
		Code:    code,
		Data:    data,
		Message: "success",
	})
}

func FailExit(r *ghttp.Request, code int, msg string) {
	r.Response.WriteJsonExit(response[string]{
		Code:    code,
		Data:    "",
		Message: msg,
	})
}
