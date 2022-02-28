package main

import (
	"log"

	"github.com/fajryhamzah/worclue/src/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server.InitServer()
}