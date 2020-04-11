package api_server

import (
	"github.com/quark_lt/cmd/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
}

/**

 */
func (api *ApiServer) NodeApi(apiRouter *mux.Router, ctl *controller.AppController) {
	nodeRouter := apiRouter.PathPrefix("/node").Subrouter()
	nodeRouter.HandleFunc("/create", ctl.CreateNode).Methods("POST")
	nodeRouter.HandleFunc("/remove/{id:[0-9]+}", ctl.RemoveNode).Methods("DELETE")
	nodeRouter.HandleFunc("/list", ctl.GetNodeList)
	nodeRouter.HandleFunc("/get/{id:[0-9]+}", ctl.GetNode)

}

/**

 */
func (api *ApiServer) TestApi(apiRouter *mux.Router, ctl *controller.AppController) {
	testRouter := apiRouter.PathPrefix("/test").Subrouter()
	testRouter.HandleFunc("/create", ctl.CreateTest).Methods("POST")
	testRouter.HandleFunc("/remove/{id:[0-9]+}", ctl.RemoveTest).Methods("DELETE")
	testRouter.HandleFunc("/list", ctl.GetTestList).Methods("GET")
	testRouter.HandleFunc("/get/{id:[0-9]+}", ctl.GetTest).Methods("GET")
	testRouter.HandleFunc("/series/{id:[0-9]+}", ctl.GetTimeSeriesData).Methods("GET")
}

/**

 */
func (api *ApiServer) UsersApi(apiRouter *mux.Router, ctl *controller.AppController) {
	userRouter := apiRouter.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/create", ctl.CreateUser).Methods("POST")
	userRouter.HandleFunc("/remove", ctl.RemoveUser).Methods("DELETE")
	userRouter.HandleFunc("/authorize", ctl.Authorize).Methods("POST")
}
func NewApiServer() *ApiServer {
	return &ApiServer{}
}

func (api *ApiServer) StartServer() {
	ctl := controller.NewAppController()

	ctl.RunMigration()
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()
	api.NodeApi(apiRouter, ctl)
	api.TestApi(apiRouter, ctl)
	api.UsersApi(apiRouter, ctl)
	log.Println("QuarkLT server started at  port: 7700 ")
	log.Println("Api server is active on http://localhost:7700/api/")
	log.Fatalln(http.ListenAndServe(":7700", r))
}
func (api *ApiServer) MigrateModels() {

}
