package apiserver_controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	models "github.com/vinogradnick/quark-lt/pkg/apiserver-models"
	"github.com/vinogradnick/quark-lt/pkg/util/config"
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

func (app *AppController) Connect(w http.ResponseWriter, r *http.Request) {

	model := models.NodeModel{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&model); err != nil {
		RespondError(w, http.StatusInternalServerError, fmt.Sprintf("%v", err))
	}

	log.Println(config.ParseJsonToString(model))
	if err := app.db.Connection.Where(models.NodeModel{Uuid: model.Uuid}).FirstOrCreate(&model).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	RespondJSON(w, http.StatusOK, "ok")
	return
}

func (app *AppController) CreateNode(w http.ResponseWriter, r *http.Request) {

	model := models.NodeModel{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&model); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	log.Println(config.ParseToString(model))
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

func (app *AppController) DeleteNode(model *models.NodeModel) {
	err := app.db.Connection.Delete(&model, "id =?", model.ID).Error
	if err != nil {
		return
	}
}
