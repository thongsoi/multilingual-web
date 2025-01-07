// main.go
package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PageData struct {
	Lang         string
	Translations map[string]map[string]string
	Content      string
}

var translations = map[string]map[string]string{
	"en": {
		"welcome":         "Welcome to our website",
		"choose_language": "Choose your language",
		"content":         "This is some sample content in English",
		"switch_to":       "Switch to",
	},
	"es": {
		"welcome":         "Bienvenido a nuestro sitio web",
		"choose_language": "Elige tu idioma",
		"content":         "Este es un contenido de ejemplo en español",
		"switch_to":       "Cambiar a",
	},
	"fr": {
		"welcome":         "Bienvenue sur notre site",
		"choose_language": "Choisissez votre langue",
		"content":         "Voici un exemple de contenu en français",
		"switch_to":       "Passer à",
	},
}

func main() {
	r := mux.NewRouter()

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	r.HandleFunc("/{lang}", handleHome).Methods("GET")
	r.HandleFunc("/content/{lang}", handleContent).Methods("GET")

	// Redirect root to English version
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/en", http.StatusMovedPermanently)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lang := vars["lang"]

	// Validate language
	if _, ok := translations[lang]; !ok {
		http.Redirect(w, r, "/en", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	data := PageData{
		Lang:         lang,
		Translations: translations,
		Content:      translations[lang]["content"],
	}

	tmpl.Execute(w, data)
}

func handleContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lang := vars["lang"]

	if _, ok := translations[lang]; !ok {
		http.Error(w, "Language not supported", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/content.html"))
	data := PageData{
		Lang:         lang,
		Translations: translations,
		Content:      translations[lang]["content"],
	}

	tmpl.Execute(w, data)
}
