package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	models "github.com/quark_lt/pkg/apiserver-models"
	"github.com/quark_lt/pkg/util/config"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func (node *QuarkNode) StartHandler(w http.ResponseWriter, r *http.Request) {
	conf := models.TestModel{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&conf); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Println(config.ParseJsonToString(conf))
	logrus.Infoln("QuarkNode success to start QuarkltConfig")
	node.Config = &conf
	node.Stop()
	fmt.Fprintln(w, "ok")
	node.Start()
	return
}
func (node *QuarkNode) WorkerConfig(w http.ResponseWriter, r *http.Request) {
	log.Println(config.ParseJsonToString(node.Config))
	RespondJSON(w, http.StatusOK, node.Config.ConfigFile)
	return
}
func (node *QuarkNode) StopHandler(w http.ResponseWriter, r *http.Request) {
	node.Stop()
	RespondJSON(w, http.StatusOK, struct{ Status string }{Status: "closed"})
	return
}
func (node *QuarkNode) StatData(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, struct {
		Status string
		Uuid   string
	}{Status: "ok", Uuid: node.NodeModel.Uuid})
	return
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}
func (node *QuarkNode) ConnectMaster() {
	jsData, err := json.Marshal(node.NodeModel)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(config.ParseToString(node.NodeConfig))
	log.Println(config.ParseJsonToString(node.NodeModel))
	res, err := http.Post(node.NodeConfig.MasterHostUrl+"/api/node/connect", "application/json", bytes.NewBuffer(jsData))
	if res != nil && res.StatusCode == http.StatusOK {
		log.Println("success add Quark Node to Master server")
	} else {
		log.Println("error of connect QuarkNode to Master server")
	}
}

func (node *QuarkNode) initRouter() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/start", node.StartHandler).Methods("POST")
	muxRouter.HandleFunc("/stop", node.StopHandler).Methods("POST")
	muxRouter.HandleFunc("/stats", node.StatData).Methods("GET")
	muxRouter.HandleFunc("/", node.WorkerConfig).Methods("GET")
	return muxRouter
}
