package quark_node

import (
	"github.com/quark_lt/cmd/quark_server"
	"github.com/quark_lt/internal/util/worker"
)

type QuarkNode struct {
	Id        int
	Scheduler *worker.WorkerScheduler
	Config    string
	Port      int
}

func NewQuarkNode(id int) *QuarkNode {
	return &QuarkNode{Id: id}
}

func RunQuarkNode() {
	server := quark_server.InitServer("quark-node-1", 3000, nil)
}
