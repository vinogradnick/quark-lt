package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"

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
