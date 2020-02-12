package quark_node

import (
	"github.com/fasthttp/router"
	"github.com/quark_lt/cmd/quark_server"
	"github.com/quark_lt/internal/util/config"
	"github.com/quark_lt/internal/util/worker"
)

type QuarkNodeScheduler struct {
	Nodes  map[int]*worker.Worker
	Config config.ScheduleConf
}

func NewQuarkNodeScheduler(port int) *QuarkNodeScheduler {
	quark_server.InitServer("worker_server", port, router.New())
	return &QuarkNodeScheduler{Nodes: map[int]*worker.Worker{}}
}
