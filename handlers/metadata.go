package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/psionikangel/scribd/db"
	"github.com/psionikangel/scribd/models"
)

//MetadataHandler : Receives metadata requests
func MetadataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var meta []models.Metadata
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&meta)
		if err != nil {
			panic(err)
		}
		for _, metadata := range meta {
			db.AddMetadata(metadata)
		}
		js, err := json.Marshal(meta)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(js))
		w.Header().Set("Content-type", "application/json")
		w.Write(js)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}
