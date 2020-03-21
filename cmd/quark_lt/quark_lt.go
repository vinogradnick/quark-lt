package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/quark_lt/cmd/quark_lt/api_server"
)

type Controller struct {
	chan
}
func main() {
	api_server.StartServer()
}
