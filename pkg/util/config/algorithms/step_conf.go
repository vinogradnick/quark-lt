package algorithms

import (
	"go.uber.org/zap"

	"github.com/quark_lt/pkg/validator"
)

type StepConf struct {
	Start    int32  `json:"start"`
	End      int32  `json:"end"`
	Duration string `json:"duration"`
	Step     int32  `json:"step"`
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

	if validator.PositiveValidate(s.Step) != nil {
		zap.L().Error("[StepConfig] Step value has Error value")
		errorValue++

	}
	return errorValue == 0

}
