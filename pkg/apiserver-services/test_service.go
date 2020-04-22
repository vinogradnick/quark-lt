package apiserver_services

import (
	"github.com/jinzhu/gorm"
	models "github.com/quark_lt/pkg/apiserver-models"
)

type TestService struct {
	connection *gorm.DB
}

func (ts *TestService) CreateTest(model *models.TestModel) {

}
