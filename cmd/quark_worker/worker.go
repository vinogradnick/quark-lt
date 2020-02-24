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
	ExportUrl  string
}

func NewWorker(host string) *Worker {
	return &Worker{
		MetricChan: make(chan metric.Metrics),
		Counter:    0,
		Client:     fasthttp.Client{Name: "http://localhost:3000"},
		StatusChan: make(chan bool),
		Uuid:       uuid.GenerateUuid(),
		ExportUrl:  host,
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
			http.Post(worker.ExportUrl, "application/json", bytes.NewBuffer(byteArr))
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

