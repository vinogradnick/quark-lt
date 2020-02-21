package config


type QuarkLTConfig struct {
	Name      string        `yaml:"name"`
	SiteSetup SiteSetupConf `yaml:"site-setup"`
}