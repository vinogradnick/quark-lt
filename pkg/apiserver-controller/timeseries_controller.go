package apiserver_controller

import (
	"fmt"
	"github.com/gorilla/mux"
	models "github.com/quark_lt/pkg/apiserver-models"

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
		ts := app.iflx.QueryDb(fmt.Sprintf(`select * from "%s" where target_server='%s' `, "gun-metrics", model.Target))
		//res := app.iflx.QueryDb(fmt.Sprintf(`select * from "%s" where target_server=%s AND time >= %s AND time <%s`, table, target, startTime, endTime))
		if ts != nil {
			RespondJSON(w, http.StatusOK, ts)
		} else {
			RespondError(w, http.StatusBadGateway, "error query influx")
		}
	}

}
