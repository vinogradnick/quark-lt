package models

import "github.com/jinzhu/gorm"

type NodeModel struct {
	gorm.Model
	id    int
	Host  string
	Uuid string
}
