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

// authentication json
type authentication struct {
	Authorized bool   `json:"authorized"`
	Error      string `json:"error"`
}

// index route
func index(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	attr := attributes{Title: "Micro Farms Colorado"}
	renderTemplate(res, "index", &attr)
}

// about route
func about(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	attr := attributes{Title: "Micro Farms Colorado - About"}
	renderTemplate(res, "about", &attr)
}

// produce route
func produce(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	attr := attributes{Title: "Micro Farms Colorado - Produce"}
	renderTemplate(res, "produce", &attr)
}

// csa route
func csa(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	attr := attributes{Title: "Micro Farms Colorado - CSA"}
	renderTemplate(res, "csa", &attr)
}

// become route
func become(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	attr := attributes{Title: "Micro Farms Colorado - Become A Micro Farm"}
	renderTemplate(res, "become", &attr)
}

// careers route
func careers(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	attr := attributes{Title: "Micro Farms Colorado - Careers"}
	renderTemplate(res, "careers", &attr)
}

// news route
func news(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	attr := attributes{Title: "Micro Farms Colorado - News & Events"}
	renderTemplate(res, "news", &attr)
}

// contact route
func contact(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	attr := attributes{Title: "Micro Farms Colorado - Contact"}
	renderTemplate(res, "contact", &attr)
}

// auth route for oauth authentication
func auth(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	result, err := authenticate(req)
	data := authentication{result, ""}
	if err != nil {
		data.Error = err.Error()
	}
	sendJSON(res, data)
}

// login route to store token and redirect to index
func login(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	attr := attributes{Title: "Login"}
	renderTemplate(res, "login", &attr)
}

// OAuth google login route
func OAuth(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	url := getAuthURL()
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}

// OAuthCallback google login route
func OAuthCallback(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	code := req.URL.Query().Get("code")
	userToken := fetchToken(code)
	user := fetchUser(userToken)
	token := createJWT(user, userToken)
	url := "/login?token=" + token
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}
