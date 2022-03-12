package server

import (
	"net/http"

	"github.com/fajryhamzah/worclue/src/exceptions"
	"github.com/fajryhamzah/worclue/src/handlers"
)

func initiateRoutes() {
	mainHandler := http.HandlerFunc(handlers.MainHandler)
	listGameHandler := http.HandlerFunc(handlers.ListGameHandler)

	http.Handle("/game", exceptions.ErrorHandlerMiddleware(setContentType(post(mainHandler))))
	http.Handle("/", exceptions.ErrorHandlerMiddleware(setContentType(get(listGameHandler))))
}
