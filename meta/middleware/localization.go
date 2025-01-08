package middleware

import (
	"context"
	"net/http"
)

type localizationKey string

const localizationKeyKey localizationKey = "localization"

func localizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		locale := r.URL.Query().Get("locale")
		if locale == "" {
			locale = "en" // default locale
		}

		ctx := context.WithValue(r.Context(), localizationKeyKey, locale)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
