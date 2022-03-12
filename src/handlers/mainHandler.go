package handlers

import (
	"encoding/json"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	gameInput := createGameInput(r)

	handler := createHandler(gameInput)
	answer := handler.GetAnswer(gameInput)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(answer)

}
