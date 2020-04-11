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
	Uuid       string
	Name       string
	Target     string
	Algorithm  string
	ConfigFile string
	Status     bool
	StartTime  time.Time
	EndTime    time.Time
}

func (model *TestModel) ConvertYaml() {
	conf := config.QuarkLTConfig{}

	log.Fatalln(yaml.Unmarshal([]byte(model.ConfigFile), &conf))
}
