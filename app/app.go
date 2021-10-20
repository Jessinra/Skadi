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
	fileServices "gitlab.com/trivery-id/skadi/internal/file/services"
	productServices "gitlab.com/trivery-id/skadi/internal/product/services"
	userController "gitlab.com/trivery-id/skadi/internal/user/controller"
	userServices "gitlab.com/trivery-id/skadi/internal/user/services"
	"gitlab.com/trivery-id/skadi/utils/logger"
	"gitlab.com/trivery-id/skadi/utils/metadata"
)

var router = gin.Default()

func StartApplication() {
	initLogger()
	initSkadiDatabase()

	initServices()
	initServiceDependencies()
	validateServices()

	initControllers()
	validateControllers()

	initRoutes()
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

	if err := fileServices.InitServices(); err != nil {
		logger.Error("Failed to initialize fileServices", err)
		panic(fmt.Sprintf("Failed to initialize fileServices: `%+v`", err))
	}
	if err := productServices.InitServices(); err != nil {
		logger.Error("Failed to initialize productServices", err)
		panic(fmt.Sprintf("Failed to initialize productServices: `%+v`", err))
	}
	if err := userServices.InitServices(); err != nil {
		logger.Error("Failed to initialize userServices", err)
		panic(fmt.Sprintf("Failed to initialize userServices: `%+v`", err))
	}
}

func initServiceDependencies() {
	// sorted alphabetically

	if err := fileServices.InitServiceDependencies(); err != nil {
		logger.Error("Failed to initialize fileServices dependencies", err)
		panic(fmt.Sprintf("Failed to initialize fileServices dependencies: `%+v`", err))
	}
	if err := productServices.InitServiceDependencies(); err != nil {
		logger.Error("Failed to initialize productServices dependencies", err)
		panic(fmt.Sprintf("Failed to initialize productServices dependencies: `%+v`", err))
	}
	if err := userServices.InitServiceDependencies(); err != nil {
		logger.Error("Failed to initialize userServices dependencies", err)
		panic(fmt.Sprintf("Failed to initialize userServices dependencies: `%+v`", err))
	}
}

func validateServices() {
	// sorted alphabetically

	if err := fileServices.ValidateServices(); err != nil {
		logger.Error("Invalid fileServices initialization", err)
		panic(fmt.Sprintf("Invalid fileServices initialization: `%+v`", err))
	}
	if err := productServices.ValidateServices(); err != nil {
		logger.Error("Invalid productServices initialization", err)
		panic(fmt.Sprintf("Invalid productServices initialization: `%+v`", err))
	}
	if err := userServices.ValidateServices(); err != nil {
		logger.Error("Invalid userServices initialization", err)
		panic(fmt.Sprintf("Invalid userServices initialization: `%+v`", err))
	}
}

func initControllers() {
	if err := userController.InitControllers(); err != nil {
		logger.Error("Failed to initialize user controllers", err)
		panic(fmt.Sprintf("Failed to initialize user controllers: `%+v`", err))
	}

	AuthController = userController.GetAuthController()
}

func validateControllers() {
	// sorted alphabetically

	if err := userController.ValidateControllers(); AuthController == nil || err != nil {
		logger.Error("Invalid userControllers initialization", err)
		panic(fmt.Sprintf("Invalid userControllers initialization: `%+v`", err))
	}
}

func getAPPSecretName() string {
	return fmt.Sprintf("skadi-%s", os.Getenv("ENV"))
}
