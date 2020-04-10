package models

import "github.com/jinzhu/gorm"

type NodeModel struct {
	gorm.Model
	Host  string
	Uuid string
}
