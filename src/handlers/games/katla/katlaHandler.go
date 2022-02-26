package katla

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fajryhamzah/worclue/src/handlers"
)

type Katla struct{}

const CODE string = "katla"
const URL string = "https://katla.vercel.app/"

func (k Katla) GetGameCode() string {
	return CODE
}

func (k Katla) GetAnswer(input handlers.GameInput) string {
	hash := input.HashInput

	if hash == "" {
		hash = k.getHash()
	}

	return k.process(hash)
}

func (k Katla) GetGameUrl() string {
	return URL
}

func (k Katla) getHash() string {
	res, err := http.Get(k.GetGameUrl())

	if err != nil {
		fmt.Println("Gagal ginii")
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		panic("Failed to get response from " + k.GetGameUrl())
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		fmt.Println("Gagal ginii1")

		panic(err)
	}

	var response katlaData
	s := doc.Find("#__NEXT_DATA__").Text()

	json.Unmarshal([]byte(s), &response)

	return response.Props.PageProps.Hashed
}

func (k Katla) process(hash string) string {
	decoded := k.decode(hash)
	split := strings.Split(decoded, "::")

	if len(split) != 3 {
		panic(errors.New("hash should be have 3 section"))
	}

	return k.decode(split[1])
}

func (k Katla) decode(hash string) string {
	realHashCount := len(hash) - 1
	numberOfEqualSign, err := strconv.Atoi(hash[realHashCount:])

	if err != nil {
		panic(errors.New("failed to convert equal sign"))
	}

	realHash := hash[0:realHashCount]
	equalSign := strings.Repeat("=", numberOfEqualSign)

	var sign int
	var fullString string
	var charByte int

	for i := 0; i < len(realHash); i++ {
		if i%2 == 0 {
			sign = -1
		} else {
			sign = 1
		}

		charByte = int(realHash[i]) + sign
		fullString += string(rune(charByte))
	}

	decoded, err := base64.StdEncoding.DecodeString(fullString + equalSign)

	if err != nil {
		fmt.Println("Gagal aaaginii", fullString+equalSign)

		panic(err)
	}

	return string(decoded)
}
