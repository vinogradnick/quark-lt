package apiserver_services

import (
	"log"
	"src/github.com/jinzhu/gorm"
)

type ApiService interface {
	GetByID(model interface{}, id int)
	GetWhere(model interface{}, where string, value interface{})
	UpdateModel(model interface{})
	RemoveModel(model interface{})
	CreateModel(model interface{})
}

type NodeService struct {
	name string
	db   *gorm.DB
}

func NewNodeService(db *gorm.DB) *NodeService {
	return &NodeService{db: db}
}

func (n *NodeService) GetByID(model interface{}, id int) {
	if err := n.db.Find(&model, "id= ?", id).Error; err != nil {
		log.Fatalf("%s get by id error", n.name)
	}
	return
}

func (n *NodeService) GetWhere(model interface{}, where string, value interface{}) {
	if err := n.db.Find(&model, where, value).Error; err != nil {
		log.Fatalf("%s getWhere error", n.name)
	}
	return
}

func (n *NodeService) UpdateModel(model interface{}) {
	if err := n.db.Save(&model).Error; err != nil {
		log.Fatalf("%s updateModel error", n.name)
	}
	return
}

func (n *NodeService) RemoveModel(model interface{}) {
	if err := n.db.Delete(&model).Error; err != nil {
		log.Fatalf("%s updateModel error", n.name)
	}
	return
}

func (n *NodeService) CreateModel(model interface{}) {
	if err := n.db.Create(&model).Error; err != nil {
		log.Fatalf("%s CreateModel error", n.name)
	}
	return
}
