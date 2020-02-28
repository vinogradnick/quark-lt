package config

type SiteSetupConf struct {
	Schedules []ScheduleConf `yaml:"schedules"`
	Helpers   []*Helpers
}
