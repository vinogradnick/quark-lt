package apiserver_controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	models "github.com/vinogradnick/quark-lt/pkg/apiserver-models"

	"net/http"
)

type AuthModel struct {
	Username string
	Password string
}

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
	RespondJSON(w, http.StatusOK, "success create user")
	return
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
func (app *AppController) GenerateToken(w http.ResponseWriter, r *http.Request) {
	model := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&model); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	newModel := models.User{}
	defer r.Body.Close()
	if err := app.db.Connection.Find(&newModel, "username =? AND password = ?", model.Username, model.Password).Error; err != nil {
		RespondError(w, http.StatusUnauthorized, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, &struct {
		Token string
	}{Token: GenerateJWT(newModel.Username)})
}
