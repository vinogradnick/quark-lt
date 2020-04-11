package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/quark_lt/cmd/api_server"
	"gitlab.com/quark_worker/pkg/quark_logger"
)

func main() {
	quark_logger.SetupLogger(quark_logger.STDOUT_LOGGER)

	server := api_server.NewApiServer()
	server.StartServer()

}
