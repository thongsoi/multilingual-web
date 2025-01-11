package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var templates map[string]*template.Template

func initTemplates() {
	templates = make(map[string]*template.Template)
	// Load templates for each language
	templates["en"] = template.Must(template.ParseFiles("subdirectories2/templates/en/index.html"))
	templates["fr"] = template.Must(template.ParseFiles("subdirectories2/templates/fr/index.html"))
	templates["es"] = template.Must(template.ParseFiles("subdirectories2/templates/es/index.html"))
}

func renderTemplate(w http.ResponseWriter, language string) {
	tmpl, exists := templates[language]
	if !exists {
		http.NotFound(w, nil)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func languageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	language := vars["lang"]

	// Default to English if language not found
	if language != "en" && language != "fr" && language != "es" {
		language = "en"
	}

	renderTemplate(w, language)
}

func main() {
	r := mux.NewRouter()

	// Route for each language
	r.HandleFunc("/{lang}/", languageHandler)

	// Static files
	r.PathPrefix("/subdirectories2/static/").Handler(http.StripPrefix("/subdirectories2/static/", http.FileServer(http.Dir("subdirectories2/static"))))

	initTemplates()

	http.ListenAndServe(":8080", r)
}
