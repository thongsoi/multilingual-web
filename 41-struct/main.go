package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	templates = template.Must(template.ParseGlob("42/templates/*.html"))
	locales   = make(map[string]map[string]string)
)

func main() {
	// Load locale data into memory
	loadLocales()

	// Create a new router
	r := mux.NewRouter()

	// Serve static files (if any)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	r.HandleFunc("/", dynamicHandler("index")).Methods("GET")
	r.HandleFunc("/about-us", dynamicHandler("about")).Methods("GET")
	r.HandleFunc("/change-language", changeLanguageHandler).Methods("POST")

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Load locale data into memory
func loadLocales() {
	// Define locale data directly in code
	locales["en"] = map[string]string{
		"title":       "Multilingual App",
		"welcome":     "Welcome!",
		"description": "This is a simple multilingual web application.",
		"aboutTitle":  "About Us",
		"aboutBody":   "We are a company dedicated to creating multilingual applications.",
	}
	locales["es"] = map[string]string{
		"title":       "Aplicación Multilingüe",
		"welcome":     "¡Bienvenido!",
		"description": "Esta es una aplicación web multilingüe simple.",
		"aboutTitle":  "Sobre Nosotros",
		"aboutBody":   "Somos una empresa dedicada a crear aplicaciones multilingües.",
	}
	locales["fr"] = map[string]string{
		"title":       "Application Multilingue",
		"welcome":     "Bienvenue!",
		"description": "Ceci est une application web multilingue simple.",
		"aboutTitle":  "À Propos",
		"aboutBody":   "Nous sommes une entreprise dédiée à la création d'applications multilingues.",
	}
}

// Dynamic handler for different pages
func dynamicHandler(page string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		if lang == "" {
			lang = "en" // Default language
		}

		data := struct {
			Lang    string
			Page    string
			Content map[string]string
		}{
			Lang:    lang,
			Page:    page,
			Content: locales[lang],
		}

		if err := templates.ExecuteTemplate(w, page+".html", data); err != nil {
			http.NotFound(w, r)
		}
	}
}

// Change language handler (HTMX endpoint)
func changeLanguageHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.FormValue("lang")
	if lang == "" {
		lang = "en" // Default language
	}

	data := struct {
		Lang    string
		Content map[string]string
	}{
		Lang:    lang,
		Content: locales[lang],
	}

	templates.ExecuteTemplate(w, "content", data)
}
