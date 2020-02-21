package parser

import (
	"fmt"
	"github.com/quark_lt/internal/util/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
)

func LoadConfig(path string, networkFlag bool) *config.QuarkLTConfig {
	var data []byte
	if networkFlag {
		data = parseUrl(path)
	} else {
		data = readfile(path)
	}
	cfg := config.QuarkLTConfig{}
	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- t:%v\n\n", cfg)
	return &cfg
}

func parseUrl(url string) []byte {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	return body
}

func readfile(path string) []byte {
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
