package main

import (
	"log"
	"net/http"

	"github.com/thongsoi/multilingual-web/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Static file server for assets
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Routes
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/content", handlers.ContentHandler)

	// Start the server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
