package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseGlob("subdirectory/templates/*/*.html"))

func main() {
	r := mux.NewRouter()

	// Serve static files
	r.PathPrefix("/subdirectory/static/").Handler(http.StripPrefix("/subdirectory/static/", http.FileServer(http.Dir("subdirectory/static"))))

	// Home route with language subdirectory
	r.HandleFunc("/{lang}/", homeHandler).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/en/", http.StatusFound) // Default to English
	})

	// HTMX endpoint to load dynamic content
	r.HandleFunc("/load-content/{lang}", contentHandler).Methods("GET")

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lang := vars["lang"]

	// Render the template for the selected language
	err := templates.ExecuteTemplate(w, lang+"/index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func contentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lang := vars["lang"]

	// Dynamic content based on language
	content := map[string]string{
		"en": "<p>This is dynamically loaded content in English!</p>",
		"fr": "<p>Ceci est un contenu dynamique en français !</p>",
		"es": "<p>¡Este es un contenido dinámico en español!</p>",
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(content[lang]))
}
