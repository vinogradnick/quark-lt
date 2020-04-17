package apiserver_models

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
	NodeId     uint
	Algorithm  string
	ConfigFile string
	Status     string
	StartTime  time.Time
	EndTime    time.Time
}

func (model *TestModel) ConvertYaml() {
	conf := config.QuarkLTConfig{}

	log.Fatalln(yaml.Unmarshal([]byte(model.ConfigFile), &conf))
}
