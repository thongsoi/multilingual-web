package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Load locale files
	loadLocales()

	// Create a new router
	r := mux.NewRouter()

	// Serve static files (if any)
	r.PathPrefix("/41-json/static/").Handler(http.StripPrefix("/41-json/static/", http.FileServer(http.Dir("41-json/static"))))

	// Routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/change-language", changeLanguageHandler).Methods("POST")

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Load locale files into memory
func loadLocales() {
	languages := []string{"en", "es", "fr"}
	for _, lang := range languages {
		file, err := os.ReadFile("41-json/locales/" + lang + ".json")
		if err != nil {
			log.Fatalf("Error loading locale file for %s: %v", lang, err)
		}
		var data map[string]string
		if err := json.Unmarshal(file, &data); err != nil {
			log.Fatalf("Error parsing locale file for %s: %v", lang, err)
		}
		locales[lang] = data
	}
}
