package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/psionikangel/scribd/db"
	"github.com/psionikangel/scribd/models"
)

//RunHandler : Enables the creation and update of runs
func RunHandler(w http.ResponseWriter, r *http.Request) {
	var run models.Run
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&run)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
	}
	if r.Method == "POST" {
		db.NewRun(run)
	} else if r.Method == "PUT" {
		db.EndRun(run)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}
