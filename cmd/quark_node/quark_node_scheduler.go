package quark_node

import (
	"github.com/fasthttp/router"
	"github.com/quark_lt/cmd/quark_server"
	"github.com/quark_lt/cmd/quark_worker/algorithm"
)

type QuarkNodeScheduler struct {
	Nodes  map[string]*algorithm.AlgoFactory
}

func NewQuarkNodeScheduler(port int) *QuarkNodeScheduler {
	quark_server.InitServer("worker_server", port, router.New())

	return &QuarkNodeScheduler{Nodes: map[string]*algorithm.AlgoFactory{}}
}
