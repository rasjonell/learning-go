// Package main includes main http server stuff
package main

import (
	"io"
	"net/http"
)

// StringHandler struct
type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	io.WriteString(writer, sh.message)
}

func main() {
	http.Handle("/message", StringHandler{message: "Hello World"})

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic(err)
	}
}
