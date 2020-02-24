package algorithm

import (
	"github.com/quark_lt/cmd/quark_worker"
	"github.com/quark_lt/internal/util/config"
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
	Worker *quark_worker.Worker
	Config *config.ScheduleConf

}

func NewAlgoFactory(worker *quark_worker.Worker,conf *config.ScheduleConf) *AlgoFactory {
	return &AlgoFactory{Worker: worker,Config:conf}
}
func (a AlgoFactory) StartTesting()  {
	if a.Config.Validate(){
		for _,data:=range a.Config.StepLoad  {
			a.Step(data)
			a.Worker.Wg.Wait()
		}
	}

}

func (a AlgoFactory) Const(conf *config.ConstConf) {
	var i int32
	startWorkers := conf.Value
	timerData := quark_worker.DurationConvertation(conf.Duration)
	tickerDuration := time.Duration(timerData)
	for i = 0; i < startWorkers; i++ {
		a.Worker.CreateHitThread()
	}

	ticker := time.NewTicker(tickerDuration * time.Millisecond)

	for range ticker.C {
		for i = 0; i < startWorkers; i++ {
			a.Worker.CreateHitThread()
		}
	}
	ticker.Stop()
	a.Worker.Wg.Wait()
	return
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
			quark_worker.WorkerCreatePool(&a.Worker)

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
