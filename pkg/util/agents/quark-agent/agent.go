package quark_agent

import (
	"fmt"
	"github.com/quark_lt/pkg/metrics"
	"github.com/quark_lt/pkg/util/config"
	"github.com/valyala/gorpc"
	"gitlab.com/quark-node/cmd/std"
	"log"
)

type AbstractAgent interface {
	GetMetrics() interface{}
	Close()
	Run()
}

type SshAgent struct {
	Port   string               `json:"port"`
	Target string               `json:"target"`
	Config *config.SshAgentConf `json:"config"`
	Client *gorpc.Client
}

func NewSshAgent(nodeConfig *config.QuarkNodeConfig, sshConfig *config.SshAgentConf) *SshAgent {
	return &SshAgent{
		Port:   "2000",
		Target: sshConfig.Host,
		Config: sshConfig,
		Client: &gorpc.Client{
			Addr: fmt.Sprintf("%s:2000", nodeConfig.ServerConfig.Host),
		},
	}
}

/*

Запуск агента тестирования

*/
func (agent *SshAgent) Run() {
	fmt.Println(config.ParseToString(agent.Config))
	std.ExecSsh("./quark_agent", agent.Port, config.ParseToString(agent.Config))
	agent.Client.Start()
}
func (agent *SshAgent) Stop() bool {
	resp, err := agent.Client.Call("close")
	if err != nil {
		log.Fatalf("Error when sending request to server: %s", err)
	}
	if resp.(string) == "closed" {
		return true
	}
	return false
}
func (agent *SshAgent) GetMetrics() *metrics.SSHMetrics {
	resp, err := agent.Client.Call("get_metrics")
	if err != nil {
		log.Fatalf("Error when sending request to server: %s", err)
	}
	log.Println("get metrics Agents")
	mtx, _ := resp.(metrics.SSHMetrics)
	return &mtx
}
