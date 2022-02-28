package server

import (
	"net/http"

	"github.com/fajryhamzah/worclue/src/exceptions"
)

func post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			exceptions.Throw(422, "Request must be POST")
		}

		next.ServeHTTP(w, r)
	})
}

func setContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
