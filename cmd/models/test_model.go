package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/quark_lt/pkg/util/config"
	"gopkg.in/yaml.v2"
)

type TestModel struct {
	gorm.Model
	Id         int
	Uuid       string
	Name       string
	Host       string
	Algorithm  string
	ConfigFile string
	Status     bool
	StartTime  time.Time
}

func (model *TestModel) ConvertYaml() {
	conf := config.QuarkLTConfig{}

	log.Fataln(yaml.Unmarshal([]byte(model.ConfigFile), &conf))
}
