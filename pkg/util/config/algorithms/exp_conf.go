package algorithms

import (
	"log"

	"github.com/quark_lt/pkg/validator"
)

type ExpConf struct {
	Value    int32  `yaml:"value"`
	Duration string `yaml:"duration"`
}

func (e ExpConf) Validate() bool {
	var errorValue int
	if validator.PositiveValidate(e.Value) != nil {
		log.Println("err")
		errorValue++
	}
	if validator.PositiveValidate(int32(validator.DurationConvertation(e.Duration))) != nil {
		log.Println("err")
		errorValue++

	}
	return errorValue == 0
}
