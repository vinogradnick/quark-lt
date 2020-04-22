package main

import (
	"context"
	"github.com/vinogradnick/quark-lt/cmd/quark-worker/app"
	"github.com/vinogradnick/quark-lt/pkg/util/quark_logger"
	"log"
	"os"
	"sync"

	"github.com/vinogradnick/quark-lt/pkg/util/config"
)

var (
	globalCtx, cancel = context.WithCancel(context.Background())
)

/*
Parse File or Load from URL
*/
func ParseDatafile(args string) *config.QuarkLTConfig {
	if len(args) > 0 {
		return config.ParseMainConfig(config.ReadFile(args))
	}
	return config.DownloadFile("http://localhost:7777/")

}

//todo: Add log rotation to file or log.file
func main() {
	args := os.Args[1:]
	quark_logger.SetupLogger(quark_logger.STDOUT_LOGGER)

	//	fileData := flag.String("f", "", "a string")
	//	databaseUrl := flag.String("db", "http://localhost:8086", "database url to write metrics")
	//flag.Parse()

	cfg := ParseDatafile(args[1])
	log.Println(config.ParseJsonToString(cfg))
	instance := app.NewAppWorker(&sync.WaitGroup{}, cfg, args[0])
	instance.Start()
}
