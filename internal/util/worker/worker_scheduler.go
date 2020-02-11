package worker

import (
	"github.com/fasthttp/router"
	"github.com/quark_lt/cmd/quark_server"
)

type WorkerScheduler struct {
	Nodes *map[int]Worker
}

func NewWorkerScheduler(port int) *WorkerScheduler {
	quark_server.InitServer("worker_server",port,router.New())
	return &WorkerScheduler{Nodes: &map[int]Worker{}}
}
