package main

import (
	"net/http"
	"strconv"
	"text/template"
)

// Context struct
type Context struct {
	Request *http.Request
	Data    []Product
}

var htmlTemplates *template.Template

func init() {
	var err error

	htmlTemplates = template.New("all")
	htmlTemplates.Funcs(map[string]any{
		"intVal": strconv.Atoi,
	})

	htmlTemplates, err = htmlTemplates.ParseGlob("templates/*.html")
	if err != nil {
		panic(err)
	}

	http.Handle(
		"/templates/",
		http.StripPrefix("/templates/",
			http.HandlerFunc(HandleTemplateRequest),
		),
	)
}

// HandleTemplateRequest func
func HandleTemplateRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path

	if path == "" {
		path = "products.html"
	}

	t := htmlTemplates.Lookup(path)

	if t == nil {
		http.NotFound(writer, request)
	} else {
		err := t.Execute(writer, Context{request, Products})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}
