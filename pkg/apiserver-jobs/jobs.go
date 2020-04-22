package apiserver_jobs

import (
	"github.com/jinzhu/gorm"
	models "github.com/quark_lt/pkg/apiserver-models"
	"log"
)

func StoopAllTests(db *gorm.DB) bool {
	var tests []*models.TestModel
	err := db.Find(&tests).Error
	if err != nil {
		log.Print("err get list Tests StopTestJob")
		return false
	}
	for _, item := range tests {
		item.Status = "stopped"
		item.NodeId = 0
		if err := db.Save(&item).Error; err != nil {
			log.Print(err.Error())
		}
	}
	return true
}
