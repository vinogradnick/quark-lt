package config

type SiteSetupConf struct {
	Schedules []ScheduleConf `yaml:"schedules"`
	Helpers   []*Helpers
	AutoStop  *AutoStop `yaml:"auto-stop"`
}
type AutoStop struct {
	Quantile      string `yaml:"quantile"`
	ResponseLimit string `yaml:"responseLimit"`
	TestTime      string `yaml:"test-time"`
}
