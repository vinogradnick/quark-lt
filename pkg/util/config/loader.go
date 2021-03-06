package config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func ParseMainConfig(arr []byte) *QuarkLTConfig {
	cfg := QuarkLTConfig{}
	err := json.Unmarshal(arr, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}

func ParseMainConfigYaml(arr []byte) *QuarkLTConfig {
	cfg := QuarkLTConfig{}
	err := yaml.Unmarshal(arr, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}

func ParseMainQuarkConfig(data string) *QuarkLTConfig {
	cfg := QuarkLTConfig{}
	err := json.Unmarshal([]byte(data), &cfg)
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

func GetJson(url string) *WorkerConfig {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	cfg := WorkerConfig{}
	if err := decoder.Decode(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}

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
func ParseDuration(data string) time.Duration {
	d, err := time.ParseDuration(data)
	if err != nil {
		panic(err)
	}
	return d
}
