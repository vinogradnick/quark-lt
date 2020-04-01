package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/quark_lt/cmd/db"
	"github.com/quark_lt/cmd/models"
)

type AppController struct {
	db *db.DbWorker
}

func (app *AppController) CreateTest(w http.ResponseWriter, r *http.Request) {
	model := models.TestModel{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&model); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := app.db.Connection.NewRecord(&model).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	db.Create(&model)
}

func createRecord(connection *gorm.DB, value interface{}) {
	connection.NewRecord(value)
	connection.Create(value)
}

/**

 */
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
