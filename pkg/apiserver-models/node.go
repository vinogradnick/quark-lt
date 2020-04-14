package apiserver_models

import "github.com/jinzhu/gorm"

type NodeModel struct {
	gorm.Model
	Name        string
	Host        string
	DatabaseUrl string
	Uuid        string
}
