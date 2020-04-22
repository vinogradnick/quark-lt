package main

import (
	"flag"
	"github.com/jinzhu/configor"
	"github.com/quark_lt/cmd/quark-node/app"
	"github.com/quark_lt/pkg/util/config"
	"github.com/quark_lt/pkg/util/quark_logger"
)

func loadConfig() *config.QuarkNodeConfig {
	apiServerConf := config.QuarkNodeConfig{}

	fileData := flag.String("f", "node.yaml", "a string")
	flag.Parse()
	configor.Load(&apiServerConf, *fileData)
	return &apiServerConf
}

func main() {
	quark_logger.SetupLogger(quark_logger.STDOUT_LOGGER)

	server := app.NewQuarkNode(loadConfig())
	server.InitCoreServer()
}
