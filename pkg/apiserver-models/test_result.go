package apiserver_models

import (
	"fmt"
	"github.com/vinogradnick/quark-lt/pkg/util/config"
	"html/template"
	"net/http"
)

type TestResult struct {
	ServerHost  interface{}
	ConfigFile  interface{}
	RPS         interface{}
	MaxRespTime interface{}
	MinRespTime interface{}
	Algorithm   interface{}
	CpuLoad     interface{}
	MemoryUsed  interface{}
	DiskLoad    interface{}
}

func NewTestResult(model *TestModel) *TestResult {
	return &TestResult{
		ServerHost: model.Target,
		ConfigFile: config.ParseToString(config.ParseMainConfig([]byte(model.ConfigFile))),
		Algorithm:  model.Algorithm,
	}
}
func (tr *TestResult) RenderThis(w http.ResponseWriter) {
	fmt.Println(config.ParseJsonToString(tr))
	tmpl, _ := template.ParseFiles("report/report.html")
	tmpl.Execute(w, tr)
}
