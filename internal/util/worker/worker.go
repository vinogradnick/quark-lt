package worker

import (
	"github.com/quark_lt/internal/util/config"
	"github.com/shipload/cmd/ship_worker/metric"
	"github.com/valyala/fasthttp"
	"sync"
)

type Worker struct {
	MetricChan chan metric.Metrics
	Counter    int32
	Host       string
	Client     fasthttp.Client
	LoadType   string
	Wg         sync.WaitGroup
	StatusChan chan bool
}

func WorkerInit(config config.ShipLoadConfig) {

	worker := Worker{
		MetricChan: make(chan metric.Metrics),
		Counter:    0,
		Host:       config.SiteSetup.Address,
		Client:     fasthttp.Client{Name: "http://localhost:3000"},
		LoadType:   config.SiteSetup.LoadType,
		StatusChan: make(chan bool),
	}
	factory := algorithm.AlgoFactory{Worker: worker}
	schedule := config.SiteSetup.Schedule
	factory.Step(schedule.StepLoad)
}
func WorkerCreatePool(worker *Worker) {
	worker.Counter++
	worker.Wg.Add(1)
	go GreenRequest(&a.Worker)
}
