package quark_node

import (
	"github.com/fasthttp/router"
	"github.com/quark_lt/cmd/quark_server"
	"github.com/quark_lt/internal/util/agents/ssh_agent"
	"github.com/quark_lt/internal/util/config"
	"github.com/valyala/fasthttp"
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
func (node QuarkNode) StartWorker() {

}

func RunQuarkNode(id int, config config.QuarkLTConfig) {
	node := NewQuarkNode(id, config)
	node.StartSshAgent()
	node.StartServer()
	node.StartWorker()
}
