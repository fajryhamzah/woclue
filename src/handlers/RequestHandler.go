package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fajryhamzah/worclue/src/exceptions"
)

func validateRequest(gameInput GameInput) {
	if gameInput.Code == "" {
		exceptions.Throw(422, "Code must not empty")
	}

	if gameInput.DateInput == "" {
		exceptions.Throw(422, "Date input must not empty")
	}
}

func createGameInput(r *http.Request) GameInput {
	var gameInput GameInput

	err := json.NewDecoder(r.Body).Decode(&gameInput)
	if err != nil {
		exceptions.Throw(422, "Input error")
	}

	validateRequest(gameInput)

	return gameInput
}
