package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := "8080"
	router := httprouter.New()
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	router.GET("/", index)
	log.Println("Running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
