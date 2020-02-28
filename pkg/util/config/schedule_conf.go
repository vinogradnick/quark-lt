package config

import (
	"github.com/quark_lt/pkg/util/config/algorithms"
)

type ScheduleConf struct {
	Routes     []string               `yaml:"routes"`
	StepLoad   *algorithms.StepConf   `yaml:"step-load"`
	ConstLoad  *algorithms.ConstConf  `yaml:"const-load"`
	LinearLoad *algorithms.LinearConf `yaml:"line-load"`
	ExpLoad    *algorithms.ExpConf    `yaml:"exp-load"`
}

func (s ScheduleConf) Validate() bool {
	validStatus := true

	if s.StepLoad != nil {
		validStatus = s.StepLoad.Validate()
	}

	return validStatus
}
