package config

type SiteSetupConf struct {
	LoadType string `yaml:"load-type"`
	Address  string `yaml:"address"`
	Schedule *ScheduleConf
}