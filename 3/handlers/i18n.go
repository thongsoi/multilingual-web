package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
)

func GetTranslations(w http.ResponseWriter, r *http.Request) {
	lang := getLanguage(r)

	// Load translations from JSON files
	translations := make(map[string]string)
	if err := loadTranslations(lang, &translations); err != nil {
		http.Error(w, "Failed to load translations: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send translations as JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(translations); err != nil {
		http.Error(w, "Failed to encode translations", http.StatusInternalServerError)
	}
}

func getLanguage(r *http.Request) string {
	// Extract language from query parameter or headers
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		lang = r.Header.Get("Accept-Language")
		if lang == "" {
			lang = "en" // default to English
		}
	}
	return lang
}

func loadTranslations(lang string, translations *map[string]string) error {
	// Build file path based on language
	filePath := filepath.Join("translations", lang+".json")

	// Check if file exists
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return errors.New("translation file not found")
	}

	// Open and decode JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(translations); err != nil {
		return errors.New("invalid translation file format")
	}

	return nil
}
