package apiserver_controller

import (
	"github.com/quark_lt/pkg/apiserver-db"
	"github.com/quark_lt/pkg/apiserver-models"
	"github.com/quark_lt/pkg/util/config"
	"log"
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
	app.db.Connection.CreateTable(&apiserver_models.TestModel{}, &apiserver_models.NodeModel{}, &apiserver_models.User{})
}
func (app *AppController) RunNode() {

}
