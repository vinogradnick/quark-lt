package apiserver_models

import (
	"github.com/jinzhu/gorm"
	"github.com/vinogradnick/quark-lt/pkg/util/config"
	"gopkg.in/yaml.v2"
	"log"
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
	StartTime  string
	EndTime    string
}

func (model *TestModel) ConvertYaml() {
	conf := config.QuarkLTConfig{}

	log.Fatalln(yaml.Unmarshal([]byte(model.ConfigFile), &conf))
}
