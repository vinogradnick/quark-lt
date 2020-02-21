package quark_worker

import (
	"bytes"
	"encoding/json"
	"github.com/quark_lt/cmd/quark_worker/algorithm"
	"github.com/quark_lt/internal/util/config"
	"github.com/quark_lt/internal/util/uuid"
	"github.com/shipload/cmd/ship_worker/metric"
	"github.com/valyala/fasthttp"
	"net/http"
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
	Uuid       string
}

func NewWorker(config config.ScheduleConf) *Worker {

	factory := algorithm.AlgoFactory{Worker: worker}
	schedule := config.SiteSetup.Schedule
	factory.Step(schedule.StepLoad)
	return &Worker{
		MetricChan: make(chan metric.Metrics),
		Counter:    0,
		Host:       config.Routes[0],
		Client:     fasthttp.Client{Name: "http://localhost:3000"},
		StatusChan: make(chan bool),
		Uuid:       uuid.GenerateUuid(),
	}

}
func (worker Worker) CreatePool() {
	worker.Counter++
	worker.Wg.Add(1)
	go GreenRequest(&a.Worker)
}

func (worker Worker) SendMetric() {
	for {
		select {
		case data := <-worker.MetricChan:
			flex, err := json.Marshal(data)
			byteArr := []byte(string(flex))
			http.Post("localhost:300", "application/json", bytes.NewBuffer(byteArr))
			if err != nil {
				panic(err)
			}
		case <-worker.StatusChan:
			worker.Wg.Done()
			return
		}
	}
}
func (w Worker) CreateHitThread() {

}

func RunWorker(config config.ScheduleConf) {
	worker := NewWorker(config)
	go worker.SendMetric()
	factory := algorithm.NewAlgoFactory(worker)
	factory.Step()
}
