package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vinogradnick/quark-lt/pkg/apiserver-controller"
	"github.com/vinogradnick/quark-lt/pkg/util/config"
	"log"
	"net/http"
)

type ApiServer struct {
	Conf *config.ApiServerConfig
}

/**

 */
func (api *ApiServer) NodeApi(apiRouter *mux.Router, ctl *apiserver_controller.AppController) {
	nodeRouter := apiRouter.PathPrefix("/node").Subrouter()
	nodeRouter.HandleFunc("/connect", ctl.Connect).Methods("POST")
	nodeRouter.HandleFunc("/create", apiserver_controller.JwtDefender(ctl.CreateNode)).Methods("POST")
	nodeRouter.HandleFunc("/remove/{id:[0-9]+}", apiserver_controller.JwtDefender(ctl.RemoveNode)).Methods("POST")
	nodeRouter.HandleFunc("/list", apiserver_controller.JwtDefender(ctl.GetNodeList))
	nodeRouter.HandleFunc("/get/{id:[0-9]+}", apiserver_controller.JwtDefender(ctl.GetNode))

}

/**

 */
func (api *ApiServer) TestApi(apiRouter *mux.Router, ctl *apiserver_controller.AppController) {
	testRouter := apiRouter.PathPrefix("/test").Subrouter()
	testRouter.HandleFunc("/create", apiserver_controller.JwtDefender(ctl.CreateTest)).Methods("POST")
	testRouter.HandleFunc("/run/{id:[0-9]+}", apiserver_controller.JwtDefender(ctl.StartTest)).Methods("POST")
	testRouter.HandleFunc("/stop/{id:[0-9]+}", apiserver_controller.JwtDefender(ctl.StopTest)).Methods("POST")
	testRouter.HandleFunc("/remove/{id:[0-9]+}", apiserver_controller.JwtDefender(ctl.RemoveTest)).Methods("POST")
	testRouter.HandleFunc("/list", apiserver_controller.JwtDefender(ctl.GetTestList)).Methods("GET")
	testRouter.HandleFunc("/get/{id:[0-9]+}", apiserver_controller.JwtDefender(ctl.GetTest)).Methods("GET")
	testRouter.HandleFunc("/series/{id:[0-9]+}", apiserver_controller.JwtDefender(ctl.GetTimeSeriesData)).Methods("GET")
	testRouter.HandleFunc("/localstop", ctl.LocalStop).Methods("POST")

}

/**

 */
func (api *ApiServer) UsersApi(apiRouter *mux.Router, ctl *apiserver_controller.AppController) {
	userRouter := apiRouter.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/create",apiserver_controller.JwtDefender(ctl.CreateUser)).Methods("POST")
	userRouter.HandleFunc("/remove/{id:[0-9]+}", apiserver_controller.JwtDefender(ctl.RemoveUser)).Methods("POST")
	userRouter.HandleFunc("/auth", ctl.GenerateToken).Methods("POST")
	userRouter.HandleFunc("/list", apiserver_controller.JwtDefender(ctl.GetUsers)).Methods("GET")

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



	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./client/static/"))))
	// Serve index page on all unhandled routes
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./client/index.html")
	})



	log.Printf("QuarkLT server started at  port: %d ", api.Conf.Port)
	log.Printf("Api server is active on http://%s:%d/api/", api.Conf.Host, api.Conf.Port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", api.Conf.Port), r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/index.html")
}
