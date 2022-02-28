package handlers

import (
	"github.com/fajryhamzah/worclue/src/enums"
	"github.com/fajryhamzah/worclue/src/exceptions"
)

func createHandler(gameInput enums.GameInput) GameHandlerInterface {
	exceptions.Throw(422, "Game handler not found")
}
