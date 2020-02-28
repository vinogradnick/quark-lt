package quark_node

import (
	"context"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/quark_lt/cmd/quark_server"
	"github.com/quark_lt/cmd/quark_worker"
	"github.com/quark_lt/cmd/quark_worker/algorithm"
	"github.com/quark_lt/pkg/util/agents/ssh_agent"
	"github.com/quark_lt/pkg/util/config"
	"github.com/quark_lt/pkg/util/uuid"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type QuarkNode struct {
	Id         int
	Scheduler  *QuarkNodeScheduler
	Config     config.QuarkLTConfig
	Port       int
	SshAgent   *ssh_agent.SshAgent
	nodeRouter *router.Router
}

func NewQuarkNode(id int, config config.QuarkLTConfig) *QuarkNode {
	return &QuarkNode{Id: id, Config: config, Port: 30000, nodeRouter: router.New()}
}
func (node QuarkNode) StartServer() {
	quark_server.InitServer(node.Config.Name, node.Port, node.nodeRouter)
}
func (node QuarkNode) StartSshAgent() {
	helpers := node.Config.SiteSetup.Helpers
	sshConfig := helpers[0].SshAgent
	if sshConfig != nil {
		panic(sshConfig)
	}
	node.SshAgent = ssh_agent.NewSshAgent(sshConfig)
	node.nodeRouter.GET("/ssh_metrics", func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString(node.SshAgent.ReadMetric())
	})
}
func (node QuarkNode) StartWorker(ctx *context.Context) {
	worker := quark_worker.NewWorker()
	node.Scheduler = NewQuarkNodeScheduler(3000)
	for _, scheduleConfig := range node.Config.SiteSetup.Schedules {
		factory := algorithm.NewAlgoFactory(worker, scheduleConfig)
		node.Scheduler.Nodes[uuid.GenerateUuid()] = factory
		factory.StartTesting()
	}
}

func RunQuarkNode(id int, config config.QuarkLTConfig) {
	ctx := context.Background()
	ctx, err := context.WithCancel(ctx)
	if err != nil {
		zap.L().Error(fmt.Sprint("Error of creating context with cancel {0}", err))
		panic(err)
	}
	node := NewQuarkNode(id, config)
	node.StartSshAgent()
	node.StartServer()
	node.StartWorker()
}
