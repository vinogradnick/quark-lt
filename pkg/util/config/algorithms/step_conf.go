package algorithms

import (
	"github.com/quark_lt/pkg/validator"
	"go.uber.org/zap"
)

type StepConf struct {
	Start    int32  `yaml:"start"`
	End      int32  `yaml:"end"`
	Duration string `yaml:"duration"`
	Step     int32  `yaml:"step"`
}

func (s StepConf) Validate() bool {
	var errorValue int
	if validator.PositiveValidate(s.Start) != nil {
		zap.L().Error("[StepConfig]  Start value has Error value")
		errorValue++
	}
	if validator.PositiveValidate(s.End) != nil {
		zap.L().Error("[StepConfig]  End value has Error value")
		errorValue++

	}
	if validator.PositiveValidate(int32(validator.DurationConvertation(s.Duration))) != nil {
		zap.L().Error("[StepConfig] Duration value has Error value")
		errorValue++

	}
	if validator.PositiveValidate(s.Step) != nil {
		zap.L().Error("[StepConfig] Step value has Error value")
		errorValue++

	}
	return errorValue==0

}