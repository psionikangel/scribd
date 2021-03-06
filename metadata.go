package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//Metadata : A file's metadata
type Metadata struct {
	Path         string
	Filesize     int64
	LastModified time.Time
	Filename     string
	Extension    string
	Checksum     string
	RunID        string
}

//MetadataList : A list of metadata
type MetadataList struct {
	Meta []Metadata
}

//MetadataHandler : Receives metadata requests
func MetadataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var meta []Metadata
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&meta)
		if err != nil {
			panic(err)
		}
		for _, metadata := range meta {
			AddMetadata(metadata)
		}
		js, err := json.Marshal(meta)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(js))
		w.Header().Set("Content-type", "application/json")
		w.Write(js)
	} else if r.Method == "GET" {
		runid := r.URL.Query().Get("runid")
		metas := FetchMetadataPerRun(runid)
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
