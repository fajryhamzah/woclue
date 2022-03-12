package handlers

type listGameResponse struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

func transformListGame(gameList map[string]GameHandlerInterface) []listGameResponse {
	var gameListResponse []listGameResponse

	for _, handler := range gameList {
		gameListResponse = append(gameListResponse, listGameResponse{
			handler.GetName(),
			handler.GetGameCode(),
			handler.GetDescription(),
			handler.GetGameUrl(),
		})
	}

	return gameListResponse
}
