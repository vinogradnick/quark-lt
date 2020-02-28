package algorithms

import (
	"github.com/quark_lt/pkg/validator"
	"go.uber.org/zap"
)

type ExpConf struct {
	Value    int32  `yaml:"value"`
	Duration string `yaml:"duration"`
}

func (e ExpConf) Validate() bool {
	var errorValue int
	if validator.PositiveValidate(e.Value) != nil {
		zap.L().Error("[ExpConfig]  Value  has Error value")
		errorValue++
	}
	if validator.PositiveValidate(int32(validator.DurationConvertation(e.Duration))) != nil {
		zap.L().Error("[ExpConfig]  Duration value has Error value")
		errorValue++

	}
	return errorValue == 0
}
