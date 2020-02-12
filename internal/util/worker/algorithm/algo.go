package algorithm

import (
	"github.com/quark_lt/internal/util/config"
	"github.com/quark_lt/internal/util/worker"
	"github.com/shipload/cmd/worker"
	"time"
)

type AbstractAlgoFactory interface {
	Step()
	Const()
	Linear()
	Exp()
	MaxPerformance()
	Stability()
	Scalling()
	Stress()
}
type AlgoFactory struct {
	Worker worker.Worker
}

func NewAlgoFactory(worker worker.Worker) *AlgoFactory {
	return &AlgoFactory{worker: worker}
}

func (a AlgoFactory) Step(config *config.StepConf) {
	var i int32
	startWorkers := config.Start
	timerData := DurationConvertation(config.Duration)
	tickerDuration := time.Duration(timerData)
	for i = 0; i < startWorkers; i++ {
		worker.WorkerCreatePool(&a.Worker)
	}

	ticker := time.NewTicker(tickerDuration * time.Millisecond)

	for range ticker.C {
		for i = 0; i < config.Step && a.Worker.Counter < config.End; i++ {
			worker.WorkerCreatePool(&a.Worker)
		}
	}
	ticker.Stop()
	a.Worker.Wg.Wait()
	return
}

func (a AlgoFactory) Const() {
	panic("implement me")
}

func (a AlgoFactory) Linear() {
	panic("implement me")
}

func (a AlgoFactory) Exp() {
	panic("implement me")
}

func (a AlgoFactory) MaxPerformance() {
	for {
		select {
		case <-a.Worker.StatusChan:
			a.Worker.Wg.Wait()
			return
		default:
			worker.WorkerCreatePool(&a.Worker)

		}
	}

}

func (a AlgoFactory) Stability() {
	panic("implement me")
}

func (a AlgoFactory) Scalling() {
	panic("implement me")
}

func (a AlgoFactory) Stress() {
	panic("implement me")
}
