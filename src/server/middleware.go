package server

import (
	"net/http"
)

func post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			panic("request must be POST")
		}

		next.ServeHTTP(w, r)
	})
}

func errorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func(w http.ResponseWriter) {
			if r := recover(); r != nil {
				w.Write([]byte("Halo semua"))
			}
		}(w)

		next.ServeHTTP(w, r)
	})
}
