package config


type ScheduleConf struct {
	StepLoad   *[]StepConf   `yaml:"step-load"`
	ConstLoad  *[]ConstConf  `yaml:"const-load"`
	LinearLoad *[]LinearConf `yaml:"line-load"`
	ExpLoad    *[]ExpConf    `yaml:"exp-load"`
}