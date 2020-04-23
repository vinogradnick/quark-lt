package app

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/SkyrisBactera/pkill"
	apiserver_models "github.com/vinogradnick/quark-lt/pkg/apiserver-models"
	node_exec "github.com/vinogradnick/quark-lt/pkg/node-exec"

	"github.com/sirupsen/logrus"
	"github.com/vinogradnick/quark-lt/pkg/util/config"
)

type QuarkNode struct {
	sync.Mutex
	Wg         *sync.WaitGroup
	NodeConfig *config.QuarkNodeConfig

	Config    *apiserver_models.TestModel
	NodeModel *apiserver_models.NodeModel
}

func NewQuarkNode(conf *config.QuarkNodeConfig) *QuarkNode {
	return &QuarkNode{
		Wg:         &sync.WaitGroup{},
		NodeConfig: conf,

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
	node_exec.ExecPart()

	//default:
	//	go node_exec.ExecPart("./quark_worker ", "w")
	//	return

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
