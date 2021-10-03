package skadi

import (
	"os"
	"strings"

	"gitlab.com/trivery-id/skadi/external/db"
	"gitlab.com/trivery-id/skadi/external/db/postgres"
	productDomain "gitlab.com/trivery-id/skadi/internal/product/domain"
	userDomain "gitlab.com/trivery-id/skadi/internal/user/domain"
	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/logger"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(cred postgres.DBCredential) error {
	opts := []db.Option{}
	if strings.EqualFold(os.Getenv("SILENT_GORM"), "true") {
		opts = append(opts, db.WithSilentMode())
	}

	gormDB, err := postgres.NewGorm(cred, opts...)
	if err != nil {
		return err
	}

	DB = gormDB
	return nil
}

func InitMigration() error {
	logger.Info("skadi database started migration!")

	if err := DB.AutoMigrate(
		&productDomain.Product{},
		&productDomain.ProductLocation{},
		&productDomain.ProductPrice{},
		&userDomain.User{},
	); err != nil {
		logger.Error("failed migrate skadi db tables", err)
		return err
	}

	logger.Info("skadi database migrated successfully!")
	return nil
}

// VerifyInitialization is a helper method to verify skadi database initialization by querying a dummy sql query.
func VerifyInitialization() error {
	const wantValue = "hello world"

	result := struct {
		Value string
	}{}

	dbResult := DB.Raw("SELECT ? AS value", wantValue).First(&result)
	if err := dbResult.Error; err != nil {
		return err
	}

	if result.Value != wantValue {
		return errors.NewInternalServerError("wrong value")
	}

	return nil
}
