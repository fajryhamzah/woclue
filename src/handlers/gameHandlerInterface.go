package handlers

import "github.com/fajryhamzah/worclue/src/enums"

type GameHandlerInterface interface {
	GetGameCode() string
	GetAnswer(input enums.GameInput) string
	GetGameUrl() string
}
