package app

import (
	"fmt"
	"github.com/quark_lt/pkg/apiserver-controller"
	"github.com/quark_lt/pkg/util/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	Conf *config.ApiServerConfig
}

/**

 */
func (api *ApiServer) NodeApi(apiRouter *mux.Router, ctl *apiserver_controller.AppController) {
	nodeRouter := apiRouter.PathPrefix("/node").Subrouter()
	nodeRouter.HandleFunc("/create", ctl.CreateNode).Methods("POST")
	nodeRouter.HandleFunc("/remove/{id:[0-9]+}", ctl.RemoveNode).Methods("POST")
	nodeRouter.HandleFunc("/list", ctl.GetNodeList)
	nodeRouter.HandleFunc("/get/{id:[0-9]+}", ctl.GetNode)

}

/**

 */
func (api *ApiServer) TestApi(apiRouter *mux.Router, ctl *apiserver_controller.AppController) {
	testRouter := apiRouter.PathPrefix("/test").Subrouter()
	testRouter.HandleFunc("/create", ctl.CreateTest).Methods("POST")
	testRouter.HandleFunc("/run/{id:[0-9]+}", ctl.StartTest).Methods("POST")
	testRouter.HandleFunc("/remove/{id:[0-9]+}", ctl.RemoveTest).Methods("POST")
	testRouter.HandleFunc("/list", ctl.GetTestList).Methods("GET")
	testRouter.HandleFunc("/get/{id:[0-9]+}", ctl.GetTest).Methods("GET")
	testRouter.HandleFunc("/series/{id:[0-9]+}", ctl.GetTimeSeriesData).Methods("GET")
}

/**

 */
func (api *ApiServer) UsersApi(apiRouter *mux.Router, ctl *apiserver_controller.AppController) {
	userRouter := apiRouter.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/create", ctl.CreateUser).Methods("POST")
	userRouter.HandleFunc("/remove", ctl.RemoveUser).Methods("POST")
	userRouter.HandleFunc("/authorize", ctl.Authorize).Methods("POST")
}
func NewApiServer(conf *config.ApiServerConfig) *ApiServer {
	return &ApiServer{Conf: conf}
}

func (api *ApiServer) StartServer() {
	ctl := apiserver_controller.NewAppController(api.Conf)

	ctl.RunMigration()
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()
	api.NodeApi(apiRouter, ctl)
	api.TestApi(apiRouter, ctl)
	api.UsersApi(apiRouter, ctl)
	log.Printf("QuarkLT server started at  port: %d ", api.Conf.Port)
	log.Printf("Api server is active on http://%s:%d/api/", api.Conf.Host, api.Conf.Port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", api.Conf.Port), r))
}
func (api *ApiServer) MigrateModels() {

}
