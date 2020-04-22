package algorithm

import (
	"fmt"
	"github.com/vinogradnick/quark-lt/pkg/util/algorithms"
	quarkConfig "github.com/vinogradnick/quark-lt/pkg/util/config"
	"github.com/vinogradnick/quark-lt/pkg/util/validator"
	"math"
	"runtime"
	"time"
)

type RpsFunction func(rps int32)

func AlgoLoop(rps int32, run RpsFunction, stop func(), endLoop time.Duration) {
	var i int32
	tickerDuration := time.NewTicker(time.Second * 1)
	go func() {
		for range tickerDuration.C {
			for i = 0; i < int32(runtime.NumCPU()); i++ {
				run(rps)
				fmt.Println(i)
			}
		}
	}()
	time.Sleep(endLoop)
	tickerDuration.Stop()
	return
}

//Const
func Const(conf *algorithms.ConstConf, run RpsFunction) {

	timerData := validator.DurationConvertation(conf.Duration) //время добавления каждого рабочего
	AlgoLoop(conf.Value, run, nil, timerData)
}
func SelectAlgo(cfg *quarkConfig.ScheduleConf, run RpsFunction, stop func()) {
	if cfg.StepLoad != nil {
		Step(cfg.StepLoad, run, stop)
	} else {
		if cfg.ConstLoad != nil {
			Const(cfg.ConstLoad, run)
		} else {
			if cfg.ExpLoad != nil {
				Exp(cfg.ExpLoad, run, stop)
			} else {
				if cfg.Performance != nil {
					MaxPerformanceAlgo(cfg.Performance, run, stop)
				}
			}
		}
	}
}

func Step(conf *algorithms.StepConf, run RpsFunction, stop func()) {

	rpsValue := conf.Start

	timerData := quarkConfig.ParseDuration(conf.Duration) //время добавления каждого рабочего

	AlgoLoop(rpsValue, run, nil, timerData)
	ticker := time.NewTicker(timerData)

	for range ticker.C {
		fmt.Println("ST")
		rpsValue += conf.Step
		if conf.End > rpsValue {
			AlgoLoop(rpsValue, run, nil, timerData)
		} else {
			ticker.Stop()
			return
		}
	}
	return
}
func Exp(config *algorithms.ExpConf, run RpsFunction, stop func()) {
	var rpsValue int32 = 1
	counter := 1

	timerData := quarkConfig.ParseDuration(config.Duration) //время добавления каждого рабочего
	fmt.Println(timerData.Milliseconds())
	AlgoLoop(rpsValue, run, nil, timerData)
	ticker := time.NewTicker(timerData)

	for range ticker.C {
		fmt.Println("ST")
		rpsValue += int32(math.Exp(float64(counter)))
		if 100000 > rpsValue {
			AlgoLoop(rpsValue, run, nil, timerData)
			counter++
		} else {
			ticker.Stop()
			return
		}
	}
}
func MaxPerformanceAlgo(config *algorithms.MaxPerformance, run RpsFunction, stop func()) {
	if config.Status {

		var rpsValue int32 = 1000

		AlgoLoop(rpsValue, run, nil, time.Second*100)
		ticker := time.NewTicker(time.Second * 100)

		for range ticker.C {
			fmt.Println("ST")
			rpsValue += 1000
			if 1000000 > rpsValue {
				AlgoLoop(rpsValue, run, nil, time.Second*100)
			} else {
				ticker.Stop()
				return
			}
		}
	}
	return
}

//func (a AlgoFactory) Exp(conf *algorithms.ExpConf) {
//	var i int32
//	startWorkers := conf.Value
//	timerData := validator.DurationConvertation(conf.Duration)
//	tickerDuration := time.Duration(timerData)
//	for i = 0; i < startWorkers; i++ {
//		worker.CreatePool(a.Worker)
//	}
//
//	ticker := time.NewTicker(tickerDuration * time.Millisecond)
//
//	for range ticker.C {
//		for i = 0; i < startWorkers; i++ {
//			worker.CreatePool(a.Worker)
//		}
//	}
//}
//
//func (a AlgoFactory) MaxPerformance() {
//	for {
//		select {
//		case <-a.Worker.StatusChan:
//			a.Worker.Wg.Wait()
//			return
//		default:
//			worker.CreatePool(a.Worker)
//		}
//	}
//
//}
//
//func (a AlgoFactory) Stability() {
//	panic("implement me")
//}
//
//func (a AlgoFactory) Scalling() {
//	panic("implement me")
//}
//
//func (a AlgoFactory) Stress() {
//	panic("implement me")
//}
//func (a *AlgoFactory) RunServer(gbCtx context.Context) {
//
//	rt := router.New()
//	rt.POST("/stop", func(ctx *fasthttp.RequestCtx) {
//		os.Exit(9)
//	})
//	rt.GET("/", func(ctx *fasthttp.RequestCtx) {
//		a.Worker.RpcClient.Start()
//		a.StartTesting(gbCtx)
//		ctx.WriteString("ok")
//	})
//	log.Println("server started")
//	err := fasthttp.ListenAndServe(fmt.Sprintf(":%s", a.Worker.Port), rt.Handler)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}
