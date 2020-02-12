package config


type ShipLoadConfig struct {
	Name      string        `yaml:"name"`
	SiteSetup SiteSetupConf `yaml:"site-setup"`
}