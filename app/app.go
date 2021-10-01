package app

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	skadipsql "gitlab.com/trivery-id/skadi/datasources/postgres/skadi"
	"gitlab.com/trivery-id/skadi/external/db/postgres"
	"gitlab.com/trivery-id/skadi/external/secret-manager/aws"
	"gitlab.com/trivery-id/skadi/utils/logger"
)

var router = gin.Default()

func StartApplication() {
	initRoutes()
	initSkadiDatabase()

	logger.Info("about to start the application...")
	if err := router.Run(":5000"); err != nil {
		panic(err.Error())
	}
}

func initSkadiDatabase() {
	secret := struct {
		DB struct {
			Skadi postgres.DBCredential `json:"skadi"`
		} `json:"database"`
	}{}

	secretName := getAPPSecretName()
	if err := aws.NewSecretManager().LoadSecret(secretName, &secret); err != nil {
		logger.Error("Failed to get secret", err)
	}

	// enable local database host port override using environment variable
	localDBHost := os.Getenv("DB_HOST")
	localDBPort := os.Getenv("DB_PORT")
	if localDBHost != "" && localDBPort != "" {
		secret.DB.Skadi.Host = localDBHost
		secret.DB.Skadi.Port, _ = strconv.Atoi(localDBPort)
	}

	if err := skadipsql.InitDatabase(secret.DB.Skadi); err != nil {
		logger.Error("Can't connect to skadi database", err)
		panic(fmt.Sprintf("Can't connect to skadi database: `%+v`", err))
	}

	go func() {
		_ = skadipsql.InitMigration()
	}()

	if err := skadipsql.VerifyInitialization(); err != nil {
		logger.Error("Invalid skadi database initialization", err)
		panic(fmt.Sprintf("Invalid skadi database initialization: `%+v`", err))
	}

	logger.Info("skadi database initialized successfully!")
}

func getAPPSecretName() string {
	return fmt.Sprintf("skadi-%s", os.Getenv("ENV"))
}
