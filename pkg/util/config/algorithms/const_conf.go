package algorithms

import (
	"github.com/quark_lt/pkg/validator"
)

type ConstConf struct {
	Value    int32  `yaml:"value"`
	Duration string `yaml:"duration"`
}

func (c ConstConf) Validate() bool{
	var errorValue int
	if validator.PositiveValidate(c.Value) != nil {
		errorValue++
	}
	if validator.PositiveValidate(int32(validator.DurationConvertation(c.Duration))) != nil {
		errorValue++

	}
	return errorValue==0
}