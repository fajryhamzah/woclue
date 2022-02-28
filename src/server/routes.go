package server

import (
	"net/http"

	"github.com/fajryhamzah/worclue/src/exceptions"
)

func placeholder(w http.ResponseWriter, r *http.Request) {

}

func initiateRoutes() {
	mainHandler := http.HandlerFunc(placeholder)
	http.Handle("/game", exceptions.ErrorHandlerMiddleware(setContentType(post(mainHandler))))
}
