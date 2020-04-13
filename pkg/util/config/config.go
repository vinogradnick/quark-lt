package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type QuarkLTConfig struct {
	Name       string        `json:"name"`
	ServerHost string        `json:"server-host"`
	SiteSetup  SiteSetupConf `json:"site-setup"`
}

type Helpers struct {
	SshAgent *SshAgentConf `json:"ssh-agent"`
}

//-----------------------------------------------------------------------------------------------------

type SshAgentConf struct {
	Host string `json:"host"`
	User string `json:"user"`
	Port int    `json:"port"`

	AuthMethod *AuthMethod `json:"auth-method"`
}

//-----------------------------------------------------------------------------------------------------

type AuthMethod struct {
	UserAuth *UserAuth `json:"user-auth"`
	KeyAuth  *KeyAuth  `json:"key-auth"`
}

//-----------------------------------------------------------------------------------------------------

type UserAuth struct {
	Password string `json:"password"`
}

//-----------------------------------------------------------------------------------------------------

type KeyAuth struct {
	Path string `json:"path"`
}

//Worker config -----------------------------------------

const (
	IN_MEMORY_TYPE = 1
	EXTERNAL_TYPE  = 2
	MICROVM_TYPE   = 3
)

type WorkerConfig struct {
	Config       *QuarkLTConfig `json:"config"`
	ServerConfig *ServerConfig  `json:"server_config"`
	DatabaseUrl  string         `json:"database_url"`
}
type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (sc *ServerConfig) GetString() string {
	return fmt.Sprintf("%s:%d", sc.Host, sc.Port)
}

type MicroVMConfig struct {
	Host   string
	Port   string
	Config interface{}
}

/// Node config
type QuarkNodeConfig struct {
	ServerConfig ServerConfig `json:"server_config"`
	DatabaseUrl  string       `json:"database_url"`
}

func ParseMainConfig(arr []byte) *QuarkLTConfig {
	cfg := QuarkLTConfig{}
	err := json.Unmarshal(arr, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
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

///------------------------------------------------------

func LoadConfig(path string, networkFlag bool) *QuarkLTConfig {
	var data []byte
	if networkFlag {
		//data = parseUrl(path)
	} else {
		data = ReadFile(path)
	}
	cfg := QuarkLTConfig{}
	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- t:%v\n\n", cfg)
	return &cfg
}
func ParseScheduleString(conf string) (ScheduleConf, error) {
	cfg := ScheduleConf{}
	err := yaml.Unmarshal([]byte(conf), &cfg)
	return cfg, err
}
func ParseSiteSetupString(data string) (SiteSetupConf, error) {
	fmt.Println(data)
	cfg := SiteSetupConf{}
	err := yaml.Unmarshal([]byte(data), &cfg)
	return cfg, err
}

func ParseSshConfig(data string) (SshAgentConf, error) {
	fmt.Println(data)
	cfg := SshAgentConf{}
	err := yaml.Unmarshal([]byte(data), &cfg)
	return cfg, err
}
func ParseToString(v interface{}) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
func ParseJsonToString(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
func DownloadFile(url string) *QuarkLTConfig {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return ParseMainConfig(body)

}

/*

Преобразование данных конфигурации рабочего

@data - структура в виде строки

*/
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

//func parseUrl(url string) []byte {
//
//	// Get the data
//	resp, err := http.Get(url)
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(r.Body)
//	return body
//}

func ReadFile(path string) []byte {
	file, fileErr := os.Open(path)
	if fileErr != nil {
		panic(fileErr)

	}
	responseData, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return responseData
}
