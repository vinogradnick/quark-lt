package models

import "github.com/jinzhu/gorm"

type Node struct {
	gorm.Model
	id    int
	Host  string
	Uuuid string
}
