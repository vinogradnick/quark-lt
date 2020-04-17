package config

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"log"
)

type QuarkNodeConfig struct {
	Name          string
	Uuid          string
	ServerConfig  ServerConfig `json:"server_config",yaml:"server_config"`
	DatabaseUrl   string       `json:"database_url",yaml:"database_url"`
	MasterHostUrl string       `json:"master_host_url",yaml:"master_host_url"`
}

func ParseQuarkNodeConfig(data string) *QuarkNodeConfig {
	cfg := QuarkNodeConfig{}
	err := yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}

func ParseQuarkNodeConfigFile(path string) *QuarkNodeConfig {
	cfg := QuarkNodeConfig{}
	err := json.Unmarshal(ReadFile(path), &cfg)
	if err != nil {
		return nil
	}
	return &cfg
}
