package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// attributes to render a page
type attributes struct {
	Title string
	Name  string
}

// index route
func index(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	attr := attributes{Title: "Index"}
	renderTemplate(res, "index", &attr)
}

// OAuth google login route
func OAuth(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	url := getAuthURL()
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}

// OAuthCallback google login route
func OAuthCallback(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	code := req.URL.Query().Get("code")
	user := fetchUser(code)
	fmt.Printf("user: %v", user)
	attr := attributes{Title: "Index"}
	renderTemplate(res, "index", &attr)
}
