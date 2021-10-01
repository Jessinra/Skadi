package app

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	skadipsql "gitlab.com/trivery-id/skadi/datasources/postgres/skadi"
	"gitlab.com/trivery-id/skadi/external/db/postgres"
	"gitlab.com/trivery-id/skadi/external/secret-manager/aws"
	"gitlab.com/trivery-id/skadi/graph/resolver"
	userServices "gitlab.com/trivery-id/skadi/internal/user/services"
	"gitlab.com/trivery-id/skadi/utils/logger"
	"gitlab.com/trivery-id/skadi/utils/metadata"
)

var router = gin.Default()

func StartApplication() {
	initLogger()
	initRoutes()
	initSkadiDatabase()

	initServices()
	initServiceDependencies()
	validateServices()

	resolver.InitResolvers()

	logger.Info("about to start the application...")
	if err := router.Run(":5000"); err != nil {
		panic(err.Error())
	}
}

func initLogger() {
	logger.InitLogger()
	logger.SetDefaultContextParser(metadata.LoggerContextparser{})
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

func initServices() {
	// sorted alphabetically

	if err := userServices.InitServices(); err != nil {
		logger.Error("Failed to initialize user services", err)
		panic(fmt.Sprintf("Failed to initialize user services: `%+v`", err))
	}
}

func initServiceDependencies() {
	// sorted alphabetically

	userServices.InitServiceDependencies()
}

func validateServices() {
	// sorted alphabetically

	if err := userServices.ValidateServices(); err != nil {
		logger.Error("Invalid userServices initialization", err)
		panic(fmt.Sprintf("Invalid userServices initialization: `%+v`", err))
	}
}

func getAPPSecretName() string {
	return fmt.Sprintf("skadi-%s", os.Getenv("ENV"))
}
