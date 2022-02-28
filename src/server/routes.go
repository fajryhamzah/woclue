package server

import (
	"net/http"

	"github.com/fajryhamzah/worclue/src/exceptions"
	"github.com/fajryhamzah/worclue/src/handlers"
)

func initiateRoutes() {
	mainHandler := http.HandlerFunc(handlers.MainHandler)
	http.Handle("/game", exceptions.ErrorHandlerMiddleware(setContentType(post(mainHandler))))
}
