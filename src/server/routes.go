package server

import "net/http"

func placeholder(w http.ResponseWriter, r *http.Request) {

}

func initiateRoutes() {
	mainHandler := http.HandlerFunc(placeholder)
	http.Handle("/game", errorHandler(post(mainHandler)))
}
