package api_server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quark_lt/cmd/db"
)

type ApiServer struct {
	dbConnection db.DbWorker
}

func (api *ApiServer) NodeApi(apiRouter *mux.Router) {
	nodeRouter := apiRouter.PathPrefix("/node").Subrouter()
	nodeRouter.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {

	}).Methods("POST")
	nodeRouter.HandleFunc("/remove", func(writer http.ResponseWriter, request *http.Request) {

	})
}

func (api *ApiServer) StartServer() {
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()
	api.NodeApi(apiRouter)

	log.Fatalln(http.ListenAndServe(":7700", r))
}
