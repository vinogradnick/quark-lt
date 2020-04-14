package config

type ApiServerConfig struct {
	Host           string `default:"localhost"`
	Port           int    `default:"7700"`
	InfluxUrl      string `default:"http://localhost:8086"`
	DatabaseConfig *ApiServerDatabaseConfig
}

func DefaultApiServerConfig() *ApiServerConfig {
	return &ApiServerConfig{DatabaseConfig: &ApiServerDatabaseConfig{}}
}

type ApiServerDatabaseConfig struct {
	Host         string `default:"localhost"`
	Port         string `default:"5432"`
	User         string `default:"root"`
	Password     string `default:"root"`
	DatabaseName string `default:"quarklt"`
	DatabaseType string `default:"sqlite"`
}
