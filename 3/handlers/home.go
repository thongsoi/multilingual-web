package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Get user's preferred language (e.g., from cookies or Accept-Language header)
	lang := getLanguage(r)

	// Load the appropriate template
	tmpl, err := template.ParseFiles(
		"3/templates/base.html",
		fmt.Sprintf("3/templates/home_%s.html", lang),
	)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	// Render the template with data (if any)
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		return
	}
}
