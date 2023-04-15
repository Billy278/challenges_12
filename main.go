package main

import (
	"log"

	"github.com/Billy278/challenges_12-13/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server.NewServer()
}
