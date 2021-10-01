package postgres

import (
	"fmt"

	"gitlab.com/trivery-id/skadi/external/db"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBCredential struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"name"`
}

func NewGorm(creds DBCredential, opts ...db.Option) (*gorm.DB, error) {
	options := db.ParseOptions(opts...)

	dbURL := getURL(creds)
	gormConfig := &gorm.Config{}
	if options.SilentMode {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	gormDB, err := gorm.Open(gormPostgres.Open(dbURL), gormConfig)
	if err != nil {
		return nil, err
	}

	if conn, err := gormDB.DB(); err == nil {
		conn.SetMaxOpenConns(options.MaxOpenConnection)
	}

	return gormDB, nil
}

// getURL return postgres Database URI.
func getURL(creds DBCredential) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		creds.Host,
		creds.Port,
		creds.Username,
		creds.Password,
		creds.DBName)
}
