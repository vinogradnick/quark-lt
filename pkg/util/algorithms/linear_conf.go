package algorithms

import (
	"github.com/vinogradnick/quark-lt/pkg/util/validator"

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
		errorValue++
	}
	if validator.PositiveValidate(l.End) != nil {
		errorValue++
	}
	if validator.PositiveValidate(int32(validator.DurationConvertation(l.Duration))) != nil {
		errorValue++
	}
	return errorValue == 0
}
