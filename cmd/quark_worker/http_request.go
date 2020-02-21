package quark_worker

import (
	"github.com/bradhe/stopwatch"
	"github.com/shipload/cmd/ship_worker/metric"
	"time"
)

func GreenRequest(worker *Worker) {
	for {
		select {
		case <-worker.StatusChan:
			worker.Wg.Done()
			return
		default:
			HttpRequest(worker)
		}

	}
}

func HttpRequest(worker *Worker) {
	watch := stopwatch.Start()
	statusCode, _, err := worker.Client.Get(nil, worker.Host)
	watch.Stop()
	if err != nil {
		panic(err)
	}
	if watch.Milliseconds() < time.Second*10 {
		worker.StatusChan <- true
	}
	worker.MetricChan <- *metric.NewMetrics(watch.Milliseconds(), statusCode)

}
