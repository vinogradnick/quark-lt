package algorithms

import (
	"github.com/quark_lt/pkg/validator"
	"go.uber.org/zap"
)

//-----------------------------------------------------------------------------------------------------
type LinearConf struct {
	Start    int32  `yaml:"start"`
	End      int32  `yaml:"end"`
	Duration string `yaml:"duration"`
}

func (l LinearConf) Validate() bool {
	var errorValue int
	if validator.PositiveValidate(l.Start) != nil {
		zap.L().Error("[LinearConf]  Start value has Error value")
		errorValue++
	}
	if validator.PositiveValidate(l.End) != nil {
		zap.L().Error("[LinearConf]  End value has Error value")
		errorValue++

	}
	if validator.PositiveValidate(int32(validator.DurationConvertation(l.Duration))) != nil {
		zap.L().Error("[LinearConf] Duration value has Error value")
		errorValue++

	}

	return errorValue==0
}