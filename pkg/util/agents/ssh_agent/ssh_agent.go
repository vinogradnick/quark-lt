package ssh_agent

import (
	"bytes"
	"fmt"
	"github.com/quark_lt/pkg/metrics"
	"github.com/quark_lt/pkg/util/config"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
)

type SshAgent struct {
	pipe    io.WriteCloser
	Session *ssh.Session
	Client  *ssh.Client
}

func NewSshAgent(conf *config.SshAgentConf) *SshAgent {
	if !config.ValidateSshAgentConf(conf) {
		return nil
	}
	sshConfig := &ssh.ClientConfig{
		User:            conf.User,
		Auth:            authParse(conf.AuthMethod),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", conf.Host, conf.Port), sshConfig)
	if err != nil {
		fmt.Println("tcp err")

	}

	return &SshAgent{
		Client: client,
	}
}
func (agent *SshAgent) ReadMetric() *metrics.SSHMetrics {
	session, err := agent.Client.NewSession()
	if err != nil {
		return nil
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	mem, cp := GetMemoryInfo(session, &b)
	return NewTargetMetrics(mem, cp)
}

func NewTargetMetrics(mem *metrics.MemoryInfo, cpu float64) *metrics.SSHMetrics {
	return &metrics.SSHMetrics{MemoryInfo: mem, CpuLoad: cpu}
}

func authParse(authMethod *config.AuthMethod) []ssh.AuthMethod {

	if authMethod.KeyAuth != nil {
		bytearr := config.ReadFile(authMethod.KeyAuth.Path)

		privateKey, err := ssh.ParsePrivateKey(bytearr)
		if err != nil {
			log.Fatal(err)
		}
		return []ssh.AuthMethod{ssh.PublicKeys(privateKey)}
		//todo исправить авторизацию пользователя
	} else {
		fmt.Println(authMethod.UserAuth.Password)
		return []ssh.AuthMethod{ssh.Password(authMethod.UserAuth.Password)}
	}
}
