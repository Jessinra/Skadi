//go:build tools
// +build tools

package main

import (
	_ "github.com/99designs/gqlgen"
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
