package config

import (
	"github.com/quark_lt/pkg/util/config/algorithms"
)

type ScheduleConf struct {
	Routing     []*RoadMap             `json:"routing"`
	StepLoad    *algorithms.StepConf   `json:"step-load"`
	ConstLoad   *algorithms.ConstConf  `json:"const-load"`
	LinearLoad  *algorithms.LinearConf `json:"line-load"`
	ExpLoad     *algorithms.ExpConf    `json:"exp-load"`
	Stress      *algorithms.Stress
	Performance *algorithms.MaxPerformance
}

type RoadMap struct {
	Url         string `json:"url"`
	Context     string `json:"context"`
	StatusCode  int    `json:"statusCode"`
	RequestType string `json:"requestType"`
}

func (s ScheduleConf) Validate() bool {
	validStatus := true

	if s.StepLoad != nil {
		validStatus = s.StepLoad.Validate()
	}

	return validStatus
}
