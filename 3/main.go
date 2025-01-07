package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thongsoi/multilingual-web/3/handlers"
)

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", handlers.HomeHandler)

	// Serve static files
	fs := http.FileServer(http.Dir(".3/templates"))
	r.PathPrefix("3/templates/").Handler(http.StripPrefix("3/templates/", fs))

	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
