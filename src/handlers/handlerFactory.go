package handlers

import (
	"github.com/fajryhamzah/worclue/src/enums"
	"github.com/fajryhamzah/worclue/src/exceptions"
	"github.com/fajryhamzah/worclue/src/handlers/games/katla"
)

func createHandler(gameInput enums.GameInput) GameHandlerInterface {
	if handler, ok := getGameListHandler()[gameInput.Code]; ok {
		return handler
	}

	exceptions.Throw(422, "Game handler not found")
	panic("to make linter silent")
}

func getGameListHandler() map[string]GameHandlerInterface {
	gameHandlerList := []GameHandlerInterface{
		katla.Katla{},
	}

	handlerByCodes := make(map[string]GameHandlerInterface)

	for _, handler := range gameHandlerList {
		handlerByCodes[handler.GetGameCode()] = handler
	}

	return handlerByCodes
}
