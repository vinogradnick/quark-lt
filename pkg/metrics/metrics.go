package metrics

import (
	"time"
)

type Metrics struct {
	ResponseTime time.Duration `json:"responseTime"`
	ResponseCode int           `json:"responseCode"`
}

func NewMetrics(responseCode int, responseTime time.Duration) *Metrics {
	return &Metrics{ResponseCode: responseCode, ResponseTime: responseTime}
}

type SSHMetrics struct {
	MemoryInfo *MemoryInfo `json:"memoryInfo"`
	CpuLoad    float64     `json:"cpuLoad"`
	IoLoad     float32     `json:"ioLoad"`
}
type MemoryInfo struct {
	Total  uint32
	Used   uint32
	Free   uint32
	Caches uint32
}
