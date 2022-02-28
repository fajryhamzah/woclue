package exceptions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Throw(errorCode int, message string) {
	panic(ErrorMessage{errorCode, message})
}

func ErrorHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func(w http.ResponseWriter) {
			if r := recover(); r != nil {
				exception := r.(ErrorMessage)

				fmt.Printf("[PANIC] %s : %d", exception.Message, exception.Code)

				response, _ := json.Marshal(exception)

				w.WriteHeader(exception.Code)
				w.Write(response)
			}
		}(w)

		next.ServeHTTP(w, r)
	})
}
