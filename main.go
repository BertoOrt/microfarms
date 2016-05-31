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

	// Navigation routes
	router.GET("/", index)
	router.GET("/about", about)
	router.GET("/find-our-produce", produce)
	router.GET("/csa", csa)
	router.GET("/become-a-micro-farm", become)
	router.GET("/careers", careers)
	router.GET("/micro-farms-news", news)
	router.GET("/contact", contact)

	// Oauth routes
	router.GET("/google", OAuth)
	router.GET("/login", login)
	router.GET("/auth", auth)
	router.GET("/auth/callback", OAuthCallback)

	log.Printf("Running on port: %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
