package db

import (

	"github.com/jinzhu/gorm"
)

type DbWorker struct {
	File             string
	Connection       *gorm.DB

}

func NewDbWorker() *DbWorker {
	return &DbWorker{File: "data.sqlite"}
}
func (db *DbWorker) Connect() {
	connection, err := gorm.Open("sqlite3", db.File)
	if err != nil {
		panic("failed to connect database")
	}
	db.Connection = connection
}

