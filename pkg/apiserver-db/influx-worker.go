package apiserver_db

import (
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	influxClient "github.com/influxdata/influxdb1-client/v2"
	"log"
	"time"
)

type InfluxDbWorker struct {
	InfluxConnection influxClient.Client
	DatabaseName     string
}

func NewInfluxDbWorker(url string) *InfluxDbWorker {
	c, err := influxClient.NewHTTPClient(influxClient.HTTPConfig{
		Addr: url,
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	return &InfluxDbWorker{InfluxConnection: c, DatabaseName: "quarklt"}
}

func (influxWorker *InfluxDbWorker) QueryDb(cmd string) interface{} {

	q := influxClient.Query{
		Command:  cmd,
		Database: influxWorker.DatabaseName,
	}

	response, err := influxWorker.InfluxConnection.Query(q)

	if err != nil {
		log.Println(err)
		return nil
	}

	return ConvertSeriesMap(response)
}
func (influxWorker *InfluxDbWorker) SimpleQuery(cmd string) *influxClient.Response {
	q := influxClient.Query{
		Command:  cmd,
		Database: influxWorker.DatabaseName,
	}

	response, err := influxWorker.InfluxConnection.Query(q)

	if err != nil {
		log.Println(err)
		return nil
	}
	return response
}
func ConvertSeriesMap(data *influxClient.Response) interface{} {
	//var arr []map[string]interface{}
	return data.Results

}
func CreateMetricServer(currentTime time.Time, rps int, timing time.Time, code int) map[string]interface{} {
	return map[string]interface{}{
		"time":         currentTime,
		"rps":          rps,
		"responseTime": timing.Second(),
		"responseCode": code,
	}
}
