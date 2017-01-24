package main

import (
	"encoding/json"
	"net/http"
	"time"
)

//Run : A instance when scribd was run
type Run struct {
	ID          string
	Machinename string
	Start       time.Time
	End         time.Time
	FilesCount  int64
}

//Runlist : A list of runs
type Runlist struct {
	Runs []Run
}

//RunHandler : Enables the creation and update of runs
func RunHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		runs := GetAllRuns()
		js, err := json.Marshal(runs)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
	var run Run
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&run)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
	}
	if r.Method == "POST" {
		NewRun(run)
	} else if r.Method == "PUT" {
		EndRun(run)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}
