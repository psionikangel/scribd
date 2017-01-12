package main

import (
	"net/http"

	"github.com/psionikangel/scribd/handlers"
)

func main() {
	http.HandleFunc("/run", handlers.RunHandler)
	http.HandleFunc("/metadata", handlers.MetadataHandler)
	http.ListenAndServe(":8080", nil)
}
