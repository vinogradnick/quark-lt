package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/quark_lt/cmd/quark_lt/api_server"
)

func main() {
	server := api_server.NewApiServer()
	server.StartServer()

}
