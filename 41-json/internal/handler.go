package internal

import (
	"html/template"
	"net/http"
)

var (
	templates = template.Must(template.ParseGlob("41-json/templates/*.html"))
	locales   = make(map[string]map[string]string)
)

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
