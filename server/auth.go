package server

import (
	"context"
	"net/http"
	"strings"
)

func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: validate a bearer token and add authed status to context
		authorization := r.Header.Get("Authorization")
		idToken := strings.TrimSpace(strings.Replace(authorization, "Bearer", "", 1))

		if idToken != "12345" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "authed", true)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
