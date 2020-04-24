package db_worker

import (
	"log"
	"time"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"

	"github.com/vinogradnick/quark-lt/pkg/metrics"
)

type DbWorker struct {
	Client       client.Client
	GunServer    string
	TargetServer string
}

func NewDbWorker(url string, target string) *DbWorker {
	log.Println(url)
	if len(url) > 0 {

		c, err := client.NewHTTPClient(client.HTTPConfig{
			Addr: url,
		})
		if err != nil {
			panic(err)
		}
		return &DbWorker{Client: c, TargetServer: target}
	}
	return &DbWorker{Client: nil, TargetServer: target}

}
func (worker *DbWorker) WriteMetrics(sshMetrics *metrics.SSHMetrics, targetMetrics []*metrics.Metrics) {
	flexTime := time.Now()

	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "quarklt",
		Precision: "1s",
	})
	log.Println(worker.TargetServer)
	tags := map[string]string{
		"target_server": worker.TargetServer,
	}
	var ps []*client.Point
	fln := len(targetMetrics)

	for _, rps := range targetMetrics {
		fields := map[string]interface{}{
			"rps":          fln,
			"responseTime": rps.ResponseTime.Milliseconds(),
			"responseCode": rps.ResponseCode,
		}

		pt, err := client.NewPoint(
			"gun-metrics",
			tags,
			fields,
			flexTime,
		)
		ps = append(ps, pt)
		if err != nil {
			panic(err)
		}

	}
	bp.AddPoints(ps)
	if sshMetrics != nil {
		pt, err := client.NewPoint(
			"ssh-metrics",
			tags,
			map[string]interface{}{
				"memory":        sshMetrics.MemoryInfo.Total,
				"memory-free":   sshMetrics.MemoryInfo.Free,
				"memory-cached": sshMetrics.MemoryInfo.Caches,
				"memory-used":   sshMetrics.MemoryInfo.Used,
				"cpu-load":      sshMetrics.CpuLoad,
				"disk-load":     sshMetrics.IoLoad,
			},
			flexTime,
		)
		if err != nil {
			log.Println("db_worker eerr")
		}
		bp.AddPoint(pt)
	}

	worker.Client.Write(bp)

}
