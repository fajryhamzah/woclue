package handlers

type GameHandlerInterface interface {
	GetGameCode() string
	GetAnswer(input GameInput) string
	GetGameUrl() string
}
