package ssh_agent

import (
	"bytes"
	"flag"
	"github.com/quark_lt/internal/util/config"
	"golang.org/x/crypto/ssh"
	"time"
)

type TargetMetrics struct {
	Memory  *MemoryInfo
	CpuLoad float64
	Date    time.Time
}

func NewTargetMetrics(mem *MemoryInfo, cpu float64) *TargetMetrics {
	return &TargetMetrics{date: time.Now(), Memory: mem, CpuLoad: cpu}
}

func InitConnection(config config.SshAgentConf) {
	var user = flag.String("u", "", "User name")
	// ...

	sshConfig := &ssh.ClientConfig{
		// указываем в конфиге имя пользователя
		User: *user,
		Auth: []ssh.AuthMethod{
			// а тут метод аутентификации по ключам
			ssh.PublicKeys(signer),
		},
	}

	//var (
	//	host = flag.String("h", config.Host, "Host")
	//	port = flag.String("p", config.Port, "Port")
	//)

	// звоним на сервер
	client, err := ssh.Dial("tcp", config.Host+" -p "+config.Port, sshConfig)
	session, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	readData(session)
	defer session.Close()

}
func readData(session *ssh.Session) {
	var b bytes.Buffer
	session.Stdout = &b
	for {
		data := NewTargetMetrics(GetMemoryInfo(session, &b), GetCpuAvInfo(session, &b))
		ls.PushBack(data)
	}
}
