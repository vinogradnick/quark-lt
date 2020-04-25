package apiserver_controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vinogradnick/quark-lt/pkg/util/config"
	"github.com/vinogradnick/quark-lt/pkg/util/uuid"

	"github.com/jinzhu/gorm"
	models "github.com/vinogradnick/quark-lt/pkg/apiserver-models"
)

type TestRunStruct struct {
	ID int
}

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
	RespondJSON(w, http.StatusOK, "sucess")

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
		ConfigFile: config.ParseJsonToString(model),
		Algorithm:  model.SiteSetup.Schedules[0].GetActive(),
	}
	if err := CreateRecord(app.db.Connection, &tModel); err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	newModel := models.TestModel{}
	if err := app.db.Connection.Find(&newModel, "uuid=?", tModel.Uuid).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		RespondJSON(w, http.StatusOK, &newModel)
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
func (app *AppController) StartTest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	test := models.TestModel{}
	node := models.NodeModel{}
	log.Println(id)
	err := app.db.Connection.Find(&test, "id =?", id).Error
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = app.db.Connection.First(&node).Error
	err = app.RunInNode(node, &test)
	if err == nil {
		test.StartTime = time.Now()
		test.Status = "active"
		test.NodeId = node.ID
		err = app.db.Connection.Save(&test).Error
	}
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, struct {
		Status string
	}{Status: "running"})
	return
}

func (app *AppController) LocalStop(w http.ResponseWriter, r *http.Request) {
	test := models.TestModel{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&test); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	t := models.TestModel{}
	fmt.Println(config.ParseToString(test))

	err := app.db.Connection.Find(&t, "name =?", test.Name).Error
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(config.ParseToString(t))
	t.Status = "stopped"
	if err := app.db.Connection.Save(&t).Error; err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, struct {
		Status string
	}{Status: "stopped"})
	return
}

func (app *AppController) StopTest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	test := models.TestModel{}
	node := models.NodeModel{}

	if err := app.db.Connection.Find(&test, "id =?", id).Error; err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := app.db.Connection.Find(&node, "id =?", test.NodeId).Error; err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	_ = app.StopInNode(node, test.ConfigFile)
	test.EndTime = time.Now()
	test.Status = "stopped"
	if err := app.db.Connection.Save(&test).Error; err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, struct {
		Status string
	}{Status: "stopped"})
	return
}

/**
Run Test inside node
*/
func (app *AppController) RunInNode(node models.NodeModel, cfg *models.TestModel) error {
	_, err := http.Post(fmt.Sprintf("http://%s/start", node.Host), "application/json", bytes.NewBuffer([]byte(config.ParseJsonToString(cfg))))
	if err != nil {
		return err
	} else {
		return nil
	}
}
func (app *AppController) StopInNode(node models.NodeModel, cfg string) error {
	_, err := http.Post(fmt.Sprintf("http://%s/stop", node.Host), "application/json", bytes.NewBuffer([]byte(cfg)))
	if err != nil {
		return err
	} else {
		return nil
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
