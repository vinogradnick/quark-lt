package config

import "fmt"

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (sc *ServerConfig) GetString() string {
	return fmt.Sprintf("%s:%d", sc.Host, sc.Port)
}
