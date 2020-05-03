package config

import (
	"github.com/vinogradnick/quark-lt/pkg/util/algorithms"
	"gopkg.in/yaml.v2"
)

type ScheduleConf struct {
	Routing     []*RoadMap             `json:"routing",yaml:"routing"`
	StepLoad    *algorithms.StepConf   `json:"step-load",yaml:"step-load"`
	ConstLoad   *algorithms.ConstConf  `json:"const-load",yaml:"const-load"`
	LinearLoad  *algorithms.LinearConf `json:"line-load",yaml:"exp-load"`
	ExpLoad     *algorithms.ExpConf    `json:"exp-load",yaml:"exp-load"`
	Stress      *algorithms.Stress     `json:"stress-load",yaml:"stress-load"`
	Performance *algorithms.MaxPerformance  `json:"performance-load",yaml:"performance-load"`
}

func (sc *ScheduleConf) GetActive() string {
	if sc.StepLoad != nil {
		return "Линейная"
	}
	if sc.ConstLoad != nil {
		return "Константная"
	}
	if sc.ExpLoad != nil {
		return "Экспоненциальная "
	}
	if sc.Performance != nil {
		return "Максимальная нагрузка"
	}
	return "Не определено"
}

type RoadMap struct {
	Url         string `json:"url",`
	Context     string `json:"context"`
	StatusCode  int    `json:"statusCode"`
	RequestType string `json:"requestType",yaml:"request-type"`
}

func (s ScheduleConf) Validate() bool {
	validStatus := true

	if s.StepLoad != nil {
		validStatus = s.StepLoad.Validate()
	}

	return validStatus
}
func ParseScheduleString(conf string) (ScheduleConf, error) {
	cfg := ScheduleConf{}
	err := yaml.Unmarshal([]byte(conf), &cfg)
	return cfg, err
}
