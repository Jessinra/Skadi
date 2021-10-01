package main

import (
	"github.com/joho/godotenv"
	"gitlab.com/trivery-id/skadi/app"
	"gitlab.com/trivery-id/skadi/utils/logger"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logger.Error("Error loading .env file", err)
	}

	app.StartApplication()
}
