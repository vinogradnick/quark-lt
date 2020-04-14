package apiserver_models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username    string
	Password    string
	AccessLevel int
}
