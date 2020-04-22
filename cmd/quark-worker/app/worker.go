package app

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/bradhe/stopwatch"
	"github.com/vinogradnick/quark-lt/pkg/metrics"
	"github.com/vinogradnick/quark-lt/pkg/util/config"
)

type RequestTypeStruct struct {
	Body string
}

func NewRequestTypeStruct(body string) *RequestTypeStruct {
	if len(body) > 0 {
		return &RequestTypeStruct{Body: body}
	} else {
		return nil
	}
}
func (rts *RequestTypeStruct) GetBytes() io.Reader {
	return strings.NewReader(rts.Body)
}

type QuarkWorker struct {
	MetricChan chan []*metrics.Metrics
	StatusChan chan bool
}

func NewQuarkWorker() *QuarkWorker {
	return &QuarkWorker{
		MetricChan: make(chan []*metrics.Metrics),
		StatusChan: make(chan bool),
	}
}

func (w *QuarkWorker) DoRequest(road *config.RoadMap) *metrics.Metrics {

	watch := stopwatch.Start()
	statusCode, err := http.Get(road.Url)
	watch.Stop()
	if err != nil {
		log.Fatal("error doRequest")
	}
	return metrics.NewMetrics(statusCode.StatusCode, watch.Milliseconds())
}
