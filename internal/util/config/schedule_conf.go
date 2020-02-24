package config

type ScheduleConf struct {
	Routes     []string      `yaml:"routes"`
	StepLoad   []*StepConf   `yaml:"step-load"`
	ConstLoad  []*ConstConf  `yaml:"const-load"`
	LinearLoad []*LinearConf `yaml:"line-load"`
	ExpLoad    []*ExpConf    `yaml:"exp-load"`
}

func (s ScheduleConf) Validate() bool{
	 status := true
	if s.StepLoad != nil {
		for _, g := range s.StepLoad {
			status=g.validate()
		}
		if !status{
			return status
		}
	}

	if s.ConstLoad != nil {
		for _, g := range s.ConstLoad {
			status=g.validate()
		}
		if !status{
			return status
		}
	}
	if s.ExpLoad != nil {
		for _, g := range s.ExpLoad {
			status=g.validate()
		}
		if !status{
			return status
		}
	}
	if s.LinearLoad != nil {
		for _, g := range s.LinearLoad {
			status=g.validate()
		}
		if !status{
			return status
		}
	}
	return status
}
