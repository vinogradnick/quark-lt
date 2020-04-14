package apiserver_db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/quark_lt/pkg/util/config"
)

type DbWorker struct {
	Config     *config.ApiServerDatabaseConfig
	File       string
	Connection *gorm.DB
}

func NewDbWorker(conf *config.ApiServerDatabaseConfig) *DbWorker {
	return &DbWorker{File: "data.sqlite", Config: conf}
}
func (db *DbWorker) Connect() {
	var connection *gorm.DB
	var err error
	switch db.Config.DatabaseType {

	case "postgres":
		connection, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", db.Config.Host, db.Config.Port, db.Config.User, db.Config.DatabaseName, db.Config.Password))
		break
	default:
		connection, err = gorm.Open("sqlite3", db.File)
		break
	}
	db.Connection = connection
	if err != nil {
		panic("failed to connect database")
	}

}
