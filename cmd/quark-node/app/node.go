package app

import (
	"fmt"
	"github.com/SkyrisBactera/pkill"
	apiserver_models "github.com/quark_lt/pkg/apiserver-models"
	node_exec "github.com/quark_lt/pkg/node-exec"
	"net/http"
	"os/exec"
	"runtime"
	"sync"

	"github.com/quark_lt/pkg/util/config"
	"github.com/sirupsen/logrus"
)

type QuarkNode struct {
	sync.Mutex
	Wg         *sync.WaitGroup
	NodeConfig *config.QuarkNodeConfig
	cmds       []*exec.Cmd
	Config     *apiserver_models.TestModel
	NodeModel  *apiserver_models.NodeModel
}

func NewQuarkNode(conf *config.QuarkNodeConfig) *QuarkNode {
	return &QuarkNode{
		Wg:         &sync.WaitGroup{},
		NodeConfig: conf,
		cmds:       []*exec.Cmd{},
		NodeModel: &apiserver_models.NodeModel{
			Name:        conf.Name,
			Host:        conf.ServerConfig.GetString(),
			DatabaseUrl: conf.DatabaseUrl,
			Uuid:        conf.Uuid,
		},
	}
}

func (node *QuarkNode) Start() {
	logrus.Println("Launch Quark Worker")
	switch runtime.GOOS {
	case "windows":
		go node_exec.ExecPart(".\\quark_worker ", fmt.Sprintf(`"%s"`, node.NodeConfig.DatabaseUrl))
		break
	default:
		go node_exec.ExecPart("./quark_worker ", fmt.Sprintf(`"%s"`, node.NodeConfig.DatabaseUrl))
	}

}
func (node *QuarkNode) Stop() {
	go func() {
		output, err := pkill.Pkill("quark_worker")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(output)
	}()
}

func (node *QuarkNode) InitCoreServer() {
	go node.ConnectMaster()
	logrus.Infoln("Quark Node is successfully started")
	logrus.Infoln("QuarkNodeScheduler  is successfully started")
	logrus.Infoln("Quark Node is active on http://localhost:7777")
	err := http.ListenAndServe(":7777", node.initRouter())

	if err != nil {
		logrus.Fatalln(err)
	}

}
