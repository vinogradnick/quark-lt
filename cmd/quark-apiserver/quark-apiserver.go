package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/configor"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/quark_lt/cmd/quark-apiserver/app"
	"github.com/quark_lt/pkg/util/config"
	"github.com/quark_lt/pkg/util/quark_logger"
)

func main() {
	fileData := flag.String("f", "core.yaml", "a string")
	flag.Parse()

	apiServerConf := config.ApiServerConfig{}
	configor.Load(&apiServerConf, *fileData)
	quark_logger.SetupLogger(quark_logger.STDOUT_LOGGER)
	if apiServerConf.DatabaseConfig == nil {
		apiServerConf.DatabaseConfig = config.DefaultApiServerConfig().DatabaseConfig
	}
	fmt.Println(config.ParseToString(apiServerConf))

	server := app.NewApiServer(&apiServerConf)
	server.StartServer()
}
