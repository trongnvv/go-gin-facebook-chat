package main

import (
	"log"
	"trongnv-chat/config"
	"trongnv-chat/database"
	"trongnv-chat/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.Init()
	database.Init()
	server.Init()
}
