package katla

import (
	"testing"

	"github.com/fajryhamzah/worclue/src/enums"
)

var katlaHandler Katla = Katla{}

func TestGetGameCode(t *testing.T) {
	if CODE != katlaHandler.GetGameCode() {
		t.Error("Wrong game code")
	}
}

func TestGetAnswer(t *testing.T) {
	answer := katlaHandler.GetAnswer(enums.GameInput{})

	if answer.Answer == "" {
		t.Error("Empty result ")
	}
}

func TestDecode(t *testing.T) {
	validOldAnswer := map[string]string{
		"dFGxZWV1": "parau",
		"cVGqZV51": "makan",
		"ZmW/bWJ1": "butir",
	}

	var answer string
	for key, value := range validOldAnswer {
		answer = katlaHandler.decode(key)

		if answer != value {
			t.Errorf("Expected %s got %s", value, answer)
		}
	}
}
