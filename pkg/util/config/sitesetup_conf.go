package config

type SiteSetupConf struct {
	Schedules []*ScheduleConf `json:"schedules"`
	Helpers   *Helpers        `json:"helpers"`
	AutoStop  *AutoStop       `json:"auto-stop"`
}

/**

 */
type AutoStop struct {
	Quantile      string `json:"quantile"`
	ResponseLimit string `json:"responseLimit"`
	TestTime      string `json:"test-time"`
}
