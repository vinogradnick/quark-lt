package apiserver_controller

import (
	"fmt"
	"time"

	"github.com/gorilla/mux"
	models "github.com/vinogradnick/quark-lt/pkg/apiserver-models"

	"net/http"
)

func (app *AppController) GetTimeSeriesData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	model := models.TestModel{}
	id := vars["id"]
	err := app.db.Connection.Find(&model, "id =?", id).Error
	if err != nil {
		RespondError(w, http.StatusBadGateway, err.Error())
	} else {
		startTime, _ := time.Parse("2006-01-02 15:04:05", model.StartTime)
		endTime, _ := time.Parse("2006-01-02 15:04:05", model.EndTime)
		ts := app.iflx.QueryDb(fmt.Sprintf(`select * from "%s" where target_server='%s' AND time>=%s AND time<= %s`, "gun-metrics", model.Target, startTime, endTime))
		//res := app.iflx.QueryDb(fmt.Sprintf(`select * from "%s" where target_server=%s AND time >= %s AND time <%s`, table, target, startTime, endTime))
		if ts != nil {
			RespondJSON(w, http.StatusOK, ts)
		} else {
			RespondError(w, http.StatusBadGateway, "error query influx")
		}
	}

}
