package apiserver_controller

import (
	"github.com/quark_lt/pkg/apiserver-db"
	apiserver_jobs "github.com/quark_lt/pkg/apiserver-jobs"
	"github.com/quark_lt/pkg/apiserver-models"
	"github.com/quark_lt/pkg/util/config"
	"log"
	"net/http"
)

type AppController struct {
	db   *apiserver_db.DbWorker
	iflx *apiserver_db.InfluxDbWorker
}

func NewAppController(cfg *config.ApiServerConfig) *AppController {
	return &AppController{
		db:   apiserver_db.NewDbWorker(cfg.DatabaseConfig),
		iflx: apiserver_db.NewInfluxDbWorker(cfg.InfluxUrl),
	}
}
func (app *AppController) RunMigration() {

	app.db.Connect()
	log.Println("RUN MIGRATION apiserver-models to database")
	app.db.Connection.AutoMigrate(&apiserver_models.TestModel{}, &apiserver_models.NodeModel{}, &apiserver_models.User{})
	//app.db.Connection.CreateTable(&apiserver_models.TestModel{}, &apiserver_models.NodeModel{}, &apiserver_models.User{})
}

func (app *AppController) StopTests() {
	if apiserver_jobs.StoopAllTests(app.db.Connection) {
		log.Println("sucess stop all tests")
	} else {
		log.Println("error stop all tests")
	}
}

func (app *AppController) LiveCheckNodes() {
	var nodes []*apiserver_models.NodeModel
	err := app.db.Connection.Find(&nodes).Error
	if err != nil {
		panic(err)
	}
	for _, node := range nodes {
		res := CheckNode(node)
		if !res {
			err = app.db.Connection.Delete(node, "id=?", node.ID).Error
			if err != nil {
				panic(err)
			}
		}
	}
}

func CheckNode(cfg *apiserver_models.NodeModel) bool {
	if res, err := http.Get(cfg.Host + "/stats"); err == nil && res != nil && res.StatusCode == http.StatusOK {
		return true
	}
	return false
}
