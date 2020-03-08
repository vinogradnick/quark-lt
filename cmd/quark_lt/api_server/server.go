package api_server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func UsersApi(apiRouter *mux.Router) {
	usersRouter := apiRouter.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {

	})
	usersRouter.HandleFunc("/remove", func(writer http.ResponseWriter, request *http.Request) {

	})
	usersRouter.HandleFunc("/auth", func(writer http.ResponseWriter, request *http.Request) {

	})
	usersRouter.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {

	})
}
func NodeApi(apiRouter *mux.Router) {
	nodeRouter := apiRouter.PathPrefix("/node").Subrouter()
	nodeRouter.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {

	}).Methods("POST")
	nodeRouter.HandleFunc("/remove", func(writer http.ResponseWriter, request *http.Request) {

	})
}
func ReportApi(apiRouter *mux.Router) {
	reportAPi := apiRouter.PathPrefix("/reports").Subrouter()
	reportAPi.HandleFunc("/all", func(writer http.ResponseWriter, request *http.Request) {

	})
	reportAPi.HandleFunc("/current/{id:[0-9]+}", func(writer http.ResponseWriter, request *http.Request) {

	})
	reportAPi.HandleFunc("/download/{id:[0-9]+}", func(writer http.ResponseWriter, request *http.Request) {

	})

}
func HistoryApi(apiRouter *mux.Router) {
	history := apiRouter.PathPrefix("/history").Subrouter()
	history.HandleFunc("/all", func(writer http.ResponseWriter, request *http.Request) {

	})
	history.HandleFunc("/current/{id:[0-9]+}", func(writer http.ResponseWriter, request *http.Request) {

	})
}

func StartServer() {
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()
	NodeApi(apiRouter)
	UsersApi(apiRouter)
	ReportApi(apiRouter)
	HistoryApi(apiRouter)
	log.Fatalln(http.ListenAndServe(":7700", r))
}
