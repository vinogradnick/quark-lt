package config

type ScheduleConf struct {
	Routes []string `yaml:"routes"`
	StepLoad   *[]StepConf   `yaml:"step-load"`
	ConstLoad  *[]ConstConf  `yaml:"const-load"`
	LinearLoad *[]LinearConf `yaml:"line-load"`
	ExpLoad    *[]ExpConf    `yaml:"exp-load"`
}
