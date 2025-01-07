package main

import (
	"log"
	"net/http"

	"github.com/thongsoi/multilingual-web/1/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve static files (if any)
	r.PathPrefix("/1/static/").Handler(http.StripPrefix("/1/static/", http.FileServer(http.Dir("1/static"))))

	// Routes
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/content", handlers.ContentHandler)

	// Start the server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
