package app

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/vinogradnick/quark-lt/cmd/quark-worker/app/db_worker"
	models "github.com/vinogradnick/quark-lt/pkg/apiserver-models"
	"github.com/vinogradnick/quark-lt/pkg/metrics"
	"github.com/vinogradnick/quark-lt/pkg/util/agents/ssh_agent"
	"github.com/vinogradnick/quark-lt/pkg/util/algorithm"
	"github.com/vinogradnick/quark-lt/pkg/util/config"
)

type AppWorker struct {
	wg               *sync.WaitGroup
	context          context.Context
	cancel           context.CancelFunc
	quarkWorker      *QuarkWorker
	sshAgent         *ssh_agent.SshAgent
	dbWorker         *db_worker.DbWorker
	cfg              *config.QuarkLTConfig
	workerStatConfig *config.WorkerConfig
}

func (aw *AppWorker) Start() {
	aw.wg.Add(1)

	go aw.StartExporter()
	aw.wg.Add(1)
	aw.RunSchedule()
	log.Println("QuarkWorker Successfully Completed ")
	aw.Stop()
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
	log.Println("start exporter")
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
				log.Println("start exporter")
				go aw.dbWorker.WriteMetrics(nil, data)
			}
		}
	}
}
func (aw *AppWorker) RunSchedule() {
	fmt.Println(config.ParseToString(aw.cfg))
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
	if aw.cancel != nil {
		aw.cancel()
	}

	aw.quarkWorker.StatusChan <- true
	aw.wg.Done()
	aw.SendStop()
}
func (aw *AppWorker) SendStop() {
	if aw.workerStatConfig.ServerConfig != nil {
		model := config.ParseJsonToString(models.TestModel{Name: aw.cfg.Name})
		res, err := http.Post("http://"+aw.workerStatConfig.ServerConfig.GetString()+"/localstop", "application/json", bytes.NewBufferString(model))
		if err != nil {
			log.Println(err)
		}
		if res.StatusCode == 200 {
			log.Println("success")
		}
	}
	aw.wg.Done()

}

/*
Создание инициализации агента и рабочего системы из конфигурации
*/
func NewAppWorker(wg *sync.WaitGroup, cfg *config.QuarkLTConfig, cfgWorker *config.WorkerConfig) *AppWorker {
	log.Println("create")
	var agent *config.SshAgentConf = nil
	if cfg.SiteSetup.Helpers != nil {
		agent = cfg.SiteSetup.Helpers.SshAgent
	}
	log.Println(cfg.ServerHost)
	log.Println("---------------------------------------------------------")
	return &AppWorker{
		wg:               wg,
		sshAgent:         ssh_agent.NewSshAgent(agent),
		quarkWorker:      NewQuarkWorker(),
		workerStatConfig: cfgWorker,
		dbWorker:         db_worker.NewDbWorker(cfgWorker.DatabaseUrl, cfg.ServerHost),
		cfg:              cfg,
	}
}
