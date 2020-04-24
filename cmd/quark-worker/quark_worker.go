package main

import (
	"context"
	"os"
	"sync"

	"github.com/vinogradnick/quark-lt/cmd/quark-worker/app"
	"github.com/vinogradnick/quark-lt/pkg/util/config"
	"github.com/vinogradnick/quark-lt/pkg/util/quark_logger"
)

var (
	globalCtx, cancel = context.WithCancel(context.Background())
)

/*
Parse File or Load from URL
*/
func ParseDatafile(args string) *config.WorkerConfig {
	if args != "w" {
		data := config.ReadFile(args)
		return config.ParseWorkerConfigFile(data)
	}

	return config.GetJson("http://localhost:7777/")

}

//todo: Add log rotation to file or log.file
func main() {
	args := os.Args[1:]
	quark_logger.SetupLogger(quark_logger.STDOUT_LOGGER)

	wConfig := ParseDatafile(args[0])
	instance := app.NewAppWorker(&sync.WaitGroup{}, wConfig.Config, wConfig)
	instance.Start()
}
