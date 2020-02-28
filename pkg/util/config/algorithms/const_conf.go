package algorithms

import (
	"github.com/quark_lt/pkg/validator"
	"go.uber.org/zap"
)

type ConstConf struct {
	Value    int32  `yaml:"value"`
	Duration string `yaml:"duration"`
}

func (c ConstConf) Validate() bool{
	var errorValue int
	if validator.PositiveValidate(c.Value) != nil {
		zap.L().Error("[ConstConfig]  Value  has Error value")
		errorValue++
	}
	if validator.PositiveValidate(int32(validator.DurationConvertation(c.Duration))) != nil {
		zap.L().Error("[ConstConfig]  Duration value has Error value")
		errorValue++

	}
	return errorValue==0
}