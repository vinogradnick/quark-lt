package quark_server

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"strconv"
)


func InitServer(name string, port int, router *router.Router)  {
	err := fasthttp.ListenAndServe(":"+strconv.Itoa(port), router.Handler)
	if err != nil {
		zap.Error(err)
	} else {
		zap.String("server", name+" is successfully started")
	}
}
