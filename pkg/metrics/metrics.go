package metrics

import "time"

type Metrics struct {
	ResponseTime time.Duration `json:"responseTime"`
	ResponseCode int           `json:"responseCode"`
}

func NewMetrics(responseCode int, responseTime time.Duration) *Metrics {
	return &Metrics{ResponseCode: responseCode, ResponseTime: responseTime}
}

type SSHMetrics struct {
	MemoryFree float32
	CpuLoad    float32
	IoLoad     float32
}
