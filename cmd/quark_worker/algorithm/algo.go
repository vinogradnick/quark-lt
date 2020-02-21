package algorithm

import (
	"github.com/quark_lt/cmd/quark_worker"
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
}

func NewAlgoFactory(worker *quark_worker.Worker) *AlgoFactory {
	return &AlgoFactory{worker: worker}
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
