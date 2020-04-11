package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/quark_lt/cmd/models"
	"log"
	"net/http"
)

func (app *AppController) RemoveNode(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]
	log.Println(id)
	err := app.db.Connection.Delete(models.NodeModel{}, "id =?", id).Error
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

}

func (app *AppController) CreateNode(w http.ResponseWriter, r *http.Request) {

	model := models.NodeModel{}
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

func (app *AppController) GetNodeList(w http.ResponseWriter, r *http.Request) {

	var nodes []*models.NodeModel
	err := app.db.Connection.Find(&nodes).Error
	if err != nil {

		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, &nodes)
	return
}
func (app *AppController) GetNode(w http.ResponseWriter, r *http.Request) {

	node := models.NodeModel{}
	vars := mux.Vars(r)
	id := vars["id"]
	err := app.db.Connection.Find(&node, "id =?", id).Error
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, node)
	return
}
