package handlers

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// HomeHandler renders the main page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

// ContentHandler dynamically loads the content based on language
func ContentHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		lang = "en" // Default language
	}

	// Load the respective content template
	switch lang {
	case "es":
		templates.ExecuteTemplate(w, "content_es.html", nil)
	case "fr":
		templates.ExecuteTemplate(w, "content_fr.html", nil)
	default:
		templates.ExecuteTemplate(w, "content_en.html", nil)
	}
}
