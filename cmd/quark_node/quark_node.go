package quark_node

import "github.com/quark_lt/internal/util/worker"

type QuarkNode struct {
	Id        int
	Scheduler *worker.WorkerScheduler
	Config    string
	Port      int

}
