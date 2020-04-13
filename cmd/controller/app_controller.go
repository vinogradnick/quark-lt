package controller

import (
	"github.com/quark_lt/cmd/db"
	"github.com/quark_lt/cmd/models"
	"log"
)

type AppController struct {
	db   *db.DbWorker
	iflx *db.InfluxDbWorker
}

func NewAppController() *AppController {
	return &AppController{db: db.NewDbWorker(), iflx: db.NewInfluxDbWorker("http://localhost:8086")}
}
func (app *AppController) RunMigration() {

	app.db.Connect()
	log.Println("RUN MIGRATION models to database")
	app.db.Connection.AutoMigrate(&models.TestModel{}, &models.NodeModel{}, &models.User{})
	app.db.Connection.CreateTable(&models.TestModel{}, &models.NodeModel{}, &models.User{})
}
func (app *AppController) RunNode() {
	
}
