package scheduler

import (
	"github.com/quark_lt/pkg/util/config"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type CoreScheduler struct {
	Nodes map[int]config.QuarkNodeConfig
	sync.Mutex
	Jobs *Queue
}

func NewCoreScheduler() *CoreScheduler {
	return &CoreScheduler{
		Nodes: map[int]config.QuarkNodeConfig{},
		Jobs:  Queue{}.Init(),
	}
}
func (scheduler *CoreScheduler) AddJob(cfg *config.QuarkLTConfig) {
	scheduler.Jobs.PushBack(cfg)
}

func (scheduler *CoreScheduler) DistributeJobs() {
	for _, node := range scheduler.Nodes {
		SendJob(node.ServerConfig.Host,scheduler.Jobs.PopFront())
	}
}
func SendJob(url string, ltConfig interface{}) {
	res, err := DoRequest(config.ParseToString(ltConfig), "POST", nil)
	if res.StatusCode == 200 {
		log.Println("Конфигурация отправлена успешно")
	}

}
func DoRequest(dataConv string, method string, headers map[string]string) (*http.Response, error) {
	client := http.Client{Timeout: time.Second * 10}

	request, err := http.NewRequest(method, w.target, strings.NewReader(dataConv))

	if err != nil {
		log.Fatalln(err)
	}
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	return client.Do(request)
}

//todo: Добавить проверку занятости сервера, если сервер доступен добавить запуск
