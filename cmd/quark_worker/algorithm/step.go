package algorithm

import (
	"github.com/quark_lt/cmd/quark_worker"
	"github.com/quark_lt/internal/util/config"
	"time"
)

func (a AlgoFactory) Step(config *config.StepConf) {
	var i int32
	startWorkers := config.Start
	timerData := quark_worker.DurationConvertation(config.Duration)
	tickerDuration := time.Duration(timerData)
	for i = 0; i < startWorkers; i++ {
		a.Worker.CreateHitThread()
	}

	ticker := time.NewTicker(tickerDuration * time.Millisecond)

	for range ticker.C {
		for i = 0; i < config.Step && a.Worker.Counter < config.End; i++ {
			a.Worker.CreateHitThread()
		}
	}
	ticker.Stop()
	a.Worker.Wg.Wait()
	return
}
