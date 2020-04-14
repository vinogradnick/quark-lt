package apiserver_controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	models "github.com/quark_lt/pkg/apiserver-models"

	"net/http"
)

func (app *AppController) CreateUser(w http.ResponseWriter, r *http.Request) {
	model := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&model); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := CreateRecord(app.db.Connection, &model); err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
func (app *AppController) RemoveUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	err := app.db.Connection.Delete(models.User{}, "id =?", id).Error
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

}
func (app *AppController) Authorize(w http.ResponseWriter, r *http.Request) {
	//no implementation
}
