package main

import (
	"log"
	"net/http"

	"web_app/handlers"
)

const portNumber = ":8081"

func main() {
	// ServeMux connects URL patterns to handler functions.
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/profile", handlers.Profile)
	mux.HandleFunc("/movies", handlers.Movies)
	mux.HandleFunc("/movie", handlers.MovieDetail)
	mux.HandleFunc("/api/recommendations", handlers.RecommendationsAPI)

	// FileServer sends files from static/ directly to the browser.
	staticFiles := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	log.Printf("Server is running at http://localhost%s", portNumber)
	if err := http.ListenAndServe(portNumber, mux); err != nil {
		log.Fatal(err)
	}
}
