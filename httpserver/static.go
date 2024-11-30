package main

import (
	"net/http"
	"text/template"
)

func init() {
	http.HandleFunc("/store", HandleStaticFileRequest)
}

// HandleStaticFileRequest func
func HandleStaticFileRequest(writer http.ResponseWriter, _ *http.Request) {
	templ, err := template.ParseFiles("./static/store.html")
	if err != nil {
		panic(err)
	}

	templ.Execute(writer, struct{ Data []Product }{Data: Products})
}
