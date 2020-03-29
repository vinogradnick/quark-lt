package metric_server

import (
	"github.com/fasthttp/router"
	"github.com/quark_lt/cmd/quark_server"
)

func RunMetricServer(name string) {
	router := router.New()
	r.GET("/ssh_metrics", HandleMetric)
	quark_server.InitServer(name, 32500, router)

}
func HandleMetric() {

}
