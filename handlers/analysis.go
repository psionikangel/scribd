package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/psionikangel/scribd/db"
)

//AnalysisHandler : Receives analysis requests
func AnalysisHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		runid := r.URL.Query().Get("runid")
		metas := db.FetchDuplicatesPerRun(runid)
		js, err := json.Marshal(metas)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}
