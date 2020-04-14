package config

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"log"
)

type WorkerConfig struct {
	Config       *QuarkLTConfig `json:"config"`
	ServerConfig *ServerConfig  `json:"server_config"`
	DatabaseUrl  string         `json:"database_url"`
}

func ParseWorkerConfig(data string) *WorkerConfig {
	cfg := WorkerConfig{}
	err := yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
func ParseWorkerConfigFile(arr []byte) *WorkerConfig {
	cfg := WorkerConfig{}
	err := json.Unmarshal(arr, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
