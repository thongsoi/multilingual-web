package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thongsoi/multilingual-web/meta/handlers"
)

// Locale represents a translation locale
type Locale string

const (
	English Locale = "en"
	Thai    Locale = "th"
)

// Translations holds localized messages
var Translations = map[Locale]map[string]string{}

// LoadTranslations loads locale files
func LoadTranslations() error {
	locales := []Locale{English, Thai}

	for _, loc := range locales {
		f, err := fs.Open(fmt.Sprintf("locales/%s.json", loc))
		if err != nil {
			return err
		}
		var data map[string]string
		if err := json.NewDecoder(f).Decode(&data); err != nil {
			return err
		}
		Translations[loc] = data
	}
	return nil
}

func main() {
	if err := LoadTranslations(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.Use(localizationMiddleware)
	r.HandleFunc("/", handlers.Index).Methods("GET")

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

//go:embed locales
var fs embed.FS
