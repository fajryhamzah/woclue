package server

import (
	"fmt"
	"net/http"
	"os"
)

func InitServer() {
	initiateRoutes()
	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))

	error := http.ListenAndServe(port, nil)

	if error != nil {
		fmt.Println(error)
		panic("failed to initialize server")
	}
}
