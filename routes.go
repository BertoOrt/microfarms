package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// attributes to render a page
type attributes struct {
	Title string
	Name  string
}

// index route
func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	attr := attributes{Title: "Index"}
	renderTemplate(w, "index", &attr)
}
