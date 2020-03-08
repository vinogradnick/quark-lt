package config

import (
	"github.com/quark_lt/pkg/util/config/algorithms"
)

type ScheduleConf struct {
	Routing    []*RoadMap             `yaml:"routing"`
	StepLoad   *algorithms.StepConf   `yaml:"step-load"`
	ConstLoad  *algorithms.ConstConf  `yaml:"const-load"`
	LinearLoad *algorithms.LinearConf `yaml:"line-load"`
	ExpLoad    *algorithms.ExpConf    `yaml:"exp-load"`
}

type RoadMap struct {
	Url         string `yaml:"url"`
	Context     string `yaml:"context"`
	StatusCode  int    `yaml:"statusCode"`
	RequestType string `yaml:"requestType"`
}

func (s ScheduleConf) Validate() bool {
	validStatus := true

	if s.StepLoad != nil {
		validStatus = s.StepLoad.Validate()
	}

	return validStatus
}
