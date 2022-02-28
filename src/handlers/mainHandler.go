package handlers

import "net/http"

func MainHandler(w http.ResponseWriter, r *http.Request) {
	gameInput := createGameInput(r)

	handler := createHandler(gameInput)

	handler.GetAnswer(gameInput)
}
