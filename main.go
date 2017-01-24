package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/run", RunHandler)
	http.HandleFunc("/metadata", MetadataHandler)
	http.HandleFunc("/analysis", AnalysisHandler)
	http.ListenAndServe(":8080", nil)
}
