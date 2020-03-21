package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"src/gopkg.in/yaml.v2"

	"os"
)

type QuarkLTConfig struct {
	Name       string        `yaml:"name"`
	ServerHost string        `yaml:"server-host"`
	SiteSetup  SiteSetupConf `yaml:"site-setup"`
}

type Helpers struct {
	SshAgent *SshAgentConf `yaml:"ssh-agent"`
}

//-----------------------------------------------------------------------------------------------------

type SshAgentConf struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Port int    `yaml:"port"`

	AuthMethod *AuthMethod `yaml:"auth-method"`
}

//-----------------------------------------------------------------------------------------------------

type AuthMethod struct {
	UserAuth *UserAuth `yaml:"user-auth"`
	KeyAuth  *KeyAuth  `yaml:"key-auth"`
}

//-----------------------------------------------------------------------------------------------------

type UserAuth struct {
	Password string `yaml:"password"`
}

//-----------------------------------------------------------------------------------------------------

type KeyAuth struct {
	Path string `yaml:"path"`
}

//Worker config -----------------------------------------

const (
	IN_MEMORY_TYPE = 1
	EXTERNAL_TYPE  = 2
	MICROVM_TYPE   = 3
)

type WorkerConfig struct {
	Config       *ScheduleConf `yaml:"config",json:"config"`
	Target       string        `yaml:"target",json:"target"`
	ServerConfig *ServerConfig `yaml:"server-config",json:"server_config"`
	ExporterUrl  string        `yaml:"export-url",json:"export_url"`
	Uuid         string        `yaml:"uuid",json:"uuid"`
	Client       string        `yaml:"client",json:"client"`
	WorkerType   int           `yaml:"worker-type",json:"worker_type"`
	Platform     string        `yaml:"platform",json:"platform"`
}
type ServerConfig struct {
	Host string `yaml:"host",json:"host"`
	Port string `yaml:"port",json:"port"`
}
type MicroVMConfig struct {
	Host   string
	Port   string
	Config interface{}
}

/// Node config
type QuarkNodeConfig struct {
	ServerConfig *ServerConfig `yaml:"server-config",json:"server_config"`
	DatabaseUrl  string        `json:"database_url"`
	ExportUrl    string        `json:"export_url"`
	WorkerType   string        `json:"worker_type"`
	WorkerCount  int           `json:"worker_count"`
	*WorkerConfig
}

func ParseQuarkNodeConfig(data string) *QuarkNodeConfig {
	cfg := QuarkNodeConfig{}
	err := yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		log.Fatal(err)
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
func ParseWorkerConfig(data string) *WorkerConfig {
	cfg := WorkerConfig{}
	err := yaml.Unmarshal([]byte(data), &cfg)
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
