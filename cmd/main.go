package main

import (
	"log"
	"os"

	"github.com/charlesoller/omni-import-microservice/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := app.Setup()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Logger.Fatal(router.Start(":"+port))
}
