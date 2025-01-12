package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var templates map[string]*template.Template

func init() {
	initTemplates()
}

func initTemplates() {
	templates = make(map[string]*template.Template)
	templates["en"] = template.Must(template.ParseFiles("subdirectories3/templates/en/index.html"))
	templates["fr"] = template.Must(template.ParseFiles("subdirectories3/templates/fr/index.html"))
	templates["es"] = template.Must(template.ParseFiles("subdirectories3/templates/es/index.html"))
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

	// Default to English if the language is invalid
	if language != "en" && language != "fr" && language != "es" {
		language = "en"
	}

	renderTemplate(w, language)
}

func dynamicQuoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	language := vars["lang"]

	// Quotes by language
	quotes := map[string][]string{
		"en": {"Keep moving forward!", "The best is yet to come.", "Success is a journey, not a destination."},
		"fr": {"Continuez à avancer!", "Le meilleur est à venir.", "Le succès est un voyage, pas une destination."},
		"es": {"¡Sigue adelante!", "Lo mejor está por venir.", "El éxito es un viaje, no un destino."},
	}

	// Default to English quotes
	langQuotes, exists := quotes[language]
	if !exists {
		langQuotes = quotes["en"]
	}

	// Randomly select a quote
	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)
	quote := langQuotes[randGenerator.Intn(len(langQuotes))]

	// Return the quote as plain text
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(quote))
}

func main() {
	r := mux.NewRouter()

	// Static file server
	r.PathPrefix("/subdirectories3/static/").Handler(http.StripPrefix("/subdirectories3/static/", http.FileServer(http.Dir("subdirectories3/static"))))

	// Main page route with language subdirectories
	r.HandleFunc("/{lang}/", languageHandler)

	// Dynamic quote endpoint
	r.HandleFunc("/{lang}/quote", dynamicQuoteHandler).Methods("GET")

	//log and run web server
	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
