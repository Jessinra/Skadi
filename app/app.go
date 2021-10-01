package app

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/trivery-id/skadi/utils/logger"
)

var router = gin.Default()

func StartApplication() {
	initRoutes()

	logger.Info("about to start the application...")
	if err := router.Run(":5000"); err != nil {
		panic(err.Error())
	}
}
