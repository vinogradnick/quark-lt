package ssh_agent

import (
	"bytes"
	"github.com/quark_lt/pkg/metrics"
	"github.com/quark_lt/pkg/parser"
	"github.com/quark_lt/pkg/util/config"
	"golang.org/x/crypto/ssh"
)

type SshAgent struct {
	session *ssh.Session
	client  *ssh.Client
}

func NewSshAgent(config *config.SshAgentConf) *SshAgent {
	sshConfig := &ssh.ClientConfig{
		User: config.User,
		Auth: authParse(config.AuthMethod),
	}

	client, err := ssh.Dial("tcp", config.Host+" -p "+config.Port, sshConfig)
	session, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	return &SshAgent{
		session: session,
		client:  client,
	}
}
func (agent SshAgent) ReadMetric() *metrics.SSHMetrics {
	var b bytes.Buffer
	agent.session.Stdout = &b
	if err != nil {
		panic(err)
	}
	return NewTargetMetrics(GetMemoryInfo(agent.session, &b), GetCpuAvInfo(agent.session, &b))
}

func NewTargetMetrics(mem *metrics.MemoryInfo, cpu float64) *metrics.SSHMetrics {
	return &metrics.SSHMetrics{MemoryInfo: mem, CpuLoad: cpu}
}

func authParse(authMethod *config.AuthMethod) []ssh.AuthMethod {
	key := []ssh.AuthMethod{}
	if authMethod.KeyAuth != nil {
		bytearr, err := parser.ReadFile(config.AuthMethod.KeyAuth.Path)
		if err != nil {
			panic(err)
		}
		privateKey, err = ssh.ParseRawPrivateKey(bytearr)
		key = []ssh.AuthMethod{privateKey}
	} else {
		key = []ssh.AuthMethod{ssh.Password(authMethod.UserAuth.Password)}
	}
	return key
}
