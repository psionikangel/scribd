package main

import (
	"encoding/json"
	"net/http"
)

//AnalysisHandler : Receives analysis requests
func AnalysisHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		runid := r.URL.Query().Get("runid")
		metas := FetchDuplicatesPerRun(runid)
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
