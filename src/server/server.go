package server

import (
	"fmt"
	"net/http"
	"os"
)

func InitServer() {
	initiateRoutes()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if port == ":" {
		port += "8000"
	}

	print("Start server at port " + port)
	error := http.ListenAndServe(port, nil)

	if error != nil {
		panic("failed to initialize server")
	}
}
