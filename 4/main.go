package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	templates = template.Must(template.ParseGlob("4/templates/*.html"))
	locales   = make(map[string]map[string]string)
)

func main() {
	// Load locale files
	loadLocales()

	// Create a new router
	r := mux.NewRouter()

	// Serve static files (if any)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

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
		file, err := os.ReadFile("4/locales/" + lang + ".json")
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

// Home handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")
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

	templates.ExecuteTemplate(w, "index.html", data)
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
