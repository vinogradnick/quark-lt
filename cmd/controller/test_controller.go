package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/quark_lt/pkg/util/config"
	"github.com/quark_lt/pkg/util/uuid"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/quark_lt/cmd/models"
)

/*


-----------------------------------------------------------------------------------------------------------------------------------------------------------

*/

func (app *AppController) RemoveTest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]
	err := app.db.Connection.Delete(models.TestModel{}, "id =?", id).Error
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

}

func (app *AppController) CreateTest(w http.ResponseWriter, r *http.Request) {

	model := config.QuarkLTConfig{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&model); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	tModel := models.TestModel{
		Uuid:       uuid.GenerateUuid(),
		Name:       model.Name,
		Target:     model.ServerHost,
		ConfigFile: config.ParseToString(model),
		Algorithm:  "Алгоритма",
	}
	if err := CreateRecord(app.db.Connection, &tModel); err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *AppController) GetTestList(w http.ResponseWriter, r *http.Request) {

	var tests []*models.TestModel
	err := app.db.Connection.Find(&tests).Error
	if err != nil {

		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, &tests)
	return
}
func (app *AppController) GetTest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	model := models.TestModel{}
	id := vars["id"]
	err := app.db.Connection.Find(&model, "id =?", id).Error
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, &model)
	return
}
func (app *AppController) GetTestByCommit(w http.ResponseWriter, r *http.Request) {

	test := models.TestModel{}
	vars := mux.Vars(r)

	commitName := vars["commit"]
	err := app.db.Connection.Find(&test, "name=?", commitName).Error
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
	}
}

/*


-----------------------------------------------------------------------------------------------------------------------------------------------------------

*/

/*


-----------------------------------------------------------------------------------------------------------------------------------------------------------

*/
// Create connection in databasel
func CreateRecord(connection *gorm.DB, value interface{}) error {
	log.Println("CREATE record database")
	connection.NewRecord(value)
	return connection.Create(value).Error
}

// return write Json Response
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}
