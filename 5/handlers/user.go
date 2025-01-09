package handlers

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	loc := r.Context().Value(localizationKey).(string)
	t := Translations[loc]

	htmx.Response(w, r, http.StatusOK, map[string]string{
		"title": t["title"],
		"hello": t["hello"],
	})
}
