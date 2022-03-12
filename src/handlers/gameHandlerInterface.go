package handlers

import "github.com/fajryhamzah/worclue/src/enums"

type GameHandlerInterface interface {
	GetGameCode() string
	GetAnswer(input enums.GameInput) enums.GameAnswer
	GetGameUrl() string
	GetName() string
	GetDescription() string
}
