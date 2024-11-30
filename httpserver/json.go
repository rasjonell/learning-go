package main

import (
	"encoding/json"
	"net/http"
)

func init() {
	http.HandleFunc("/api", HandleJSONRequest)
}

// HandleJSONRequest func
func HandleJSONRequest(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(Products)
}
