package app

import (
	"context"
	"fmt"
	"github.com/quark_lt/cmd/quark-worker/app/db_worker"
	"github.com/quark_lt/pkg/util/algorithm"

	"log"
	"os"
	"sync"

	"github.com/quark_lt/pkg/metrics"
	"github.com/quark_lt/pkg/util/agents/ssh_agent"
	"github.com/quark_lt/pkg/util/config"
)

type AppWorker struct {
	wg          *sync.WaitGroup
	context     context.Context
	cancel      context.CancelFunc
	quarkWorker *QuarkWorker
	sshAgent    *ssh_agent.SshAgent
	dbWorker    *db_worker.DbWorker
	cfg         *config.QuarkLTConfig
}

func (aw *AppWorker) Start() {
	aw.wg.Add(2)

	go aw.StartExporter()
	aw.RunSchedule()
	log.Println("QuarkWorker Successfully Completed ")
	os.Exit(9)
	aw.wg.Wait()

}

/*
Запуск пула выполнения запросов
*/
func (aw *AppWorker) StartPool(rps int32, road []*config.RoadMap) {
	var i int32
	var arr []*metrics.Metrics
	for i = 0; i < rps; i++ {
		select {
		case <-aw.quarkWorker.StatusChan:
			aw.wg.Done()
			return
		default:

			for _, request := range road {
				arr = append(arr, aw.quarkWorker.DoRequest(request))
			}

		}

	}
	if arr != nil && len(arr) > 0 {
		aw.quarkWorker.MetricChan <- arr
		log.Println("current RPS==> ", len(arr))

	}

}

/*

 Запуск агента экспортирования метрических данных

*/
func (aw *AppWorker) StartExporter() {
	for {
		select {
		case <-aw.quarkWorker.StatusChan:
			aw.wg.Done()
			aw.cancel()
			return
		case data := <-aw.quarkWorker.MetricChan:
			if aw.sshAgent != nil {
				sshData := aw.sshAgent.ReadMetric()
				go aw.dbWorker.WriteMetrics(sshData, data)
			} else {
				go aw.dbWorker.WriteMetrics(nil, data)
			}
		}
	}
}
func (aw *AppWorker) RunSchedule() {
	fmt.Println(aw.cfg.SiteSetup)
	for _, sh := range aw.cfg.SiteSetup.Schedules {

		if sh != nil {
			if sh.Validate() {
				algorithm.SelectAlgo(sh, func(rps int32) { go aw.StartPool(rps, sh.Routing) }, func() { aw.Stop() })
			} else {
				log.Fatalln("Configuration Schedule Error")
			}
		}

	}
}

/**
Остановка работы программы
*/
func (aw *AppWorker) Stop() {
	aw.cancel()
	aw.quarkWorker.StatusChan <- true
	aw.wg.Done()
}

/*
Создание инициализации агента и рабочего системы из конфигурации
*/
func NewAppWorker(wg *sync.WaitGroup, cfg *config.QuarkLTConfig, databaseUrl string) *AppWorker {
	log.Println("create")
	var agent *config.SshAgentConf = nil
	if cfg.SiteSetup.Helpers != nil {
		agent = cfg.SiteSetup.Helpers.SshAgent
	}
	return &AppWorker{
		wg:          wg,
		sshAgent:    ssh_agent.NewSshAgent(agent),
		quarkWorker: NewQuarkWorker(),
		dbWorker:    db_worker.NewDbWorker(databaseUrl, cfg.ServerHost),
		cfg:         cfg,
	}
}
