package handlers

import (
	"github.com/fajryhamzah/worclue/src/enums"
	"github.com/fajryhamzah/worclue/src/exceptions"
	"github.com/fajryhamzah/worclue/src/handlers/games/katla"
)

func createHandler(gameInput enums.GameInput) GameHandlerInterface {
	switch gameInput.Code {
	case katla.CODE:
		return katla.Katla{}
	}

	exceptions.Throw(422, "Game handler not found")
	panic("to make linter silent")
}
