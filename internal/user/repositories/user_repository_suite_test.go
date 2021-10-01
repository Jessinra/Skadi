package repositories_test

import (
	"fmt"

	"github.com/stretchr/testify/suite"
	option "gitlab.com/trivery-id/skadi/external/db"
	"gitlab.com/trivery-id/skadi/external/db/postgres"
	"gitlab.com/trivery-id/skadi/internal/user/domain"
	"gitlab.com/trivery-id/skadi/internal/user/repositories"
	"gorm.io/gorm"
)

type UserRepositorySuite struct {
	suite.Suite
	Credential postgres.DBCredential
	TestDB     *gorm.DB
	EmptyDB    *gorm.DB

	UserRepository      *repositories.UserRepository
	EmptyUserRepository *repositories.UserRepository
}

func NewUserRepositorySuite(cred postgres.DBCredential) *UserRepositorySuite {
	return &UserRepositorySuite{
		Credential: cred,
	}
}

// SetupSuite run once before all test, implement testify/suite's SetupAllSuite interface.
func (s *UserRepositorySuite) SetupSuite() {
	s.T().Log("UserRepositorySuite SetupSuite")

	s.setupTestDB()
	s.setupEmptyDB()

	s.UserRepository = &repositories.UserRepository{DB: s.TestDB}
	s.EmptyUserRepository = &repositories.UserRepository{DB: s.EmptyDB}
}

func (s *UserRepositorySuite) setupTestDB() {
	gormDB, err := postgres.NewGorm(s.Credential, option.WithSilentMode())
	if err != nil {
		s.T().Fatalf("Failed to connect database: %v", err)
	}

	if db, err := gormDB.DB(); err != nil && db.Ping() != nil {
		s.T().Fatalf("Failed to connect database: %v", err)
	}

	if err := gormDB.AutoMigrate(
		&domain.User{},
	); err != nil {
		s.T().Fatal(err.Error())
	}

	s.TestDB = gormDB
}

func (s *UserRepositorySuite) setupEmptyDB() {
	const emptyDBName = "database_empty"

	s.TestDB.Exec(fmt.Sprintf("CREATE DATABASE %s", emptyDBName))
	emptyDBCredential := s.Credential
	emptyDBCredential.DBName = emptyDBName

	gormDB, err := postgres.NewGorm(emptyDBCredential, option.WithSilentMode())
	if err != nil {
		s.T().Fatalf("Failed to connect to empty database: %v", err)
	}

	if db, err := gormDB.DB(); err != nil && db.Ping() != nil {
		s.T().Fatalf("Failed to connect to empty database: %v", err)
	}

	s.EmptyDB = gormDB
}

// TearDownSuite run once after all test, implement testify/suite's TearDownAllSuite interface.
func (s *UserRepositorySuite) TearDownSuite() {
	s.T().Log("UserRepositorySuite TearDownSuite")
	if db, _ := s.TestDB.DB(); db != nil {
		db.Close()
	}
	if db, _ := s.EmptyDB.DB(); db != nil {
		db.Close()
	}
}

// SetupTest run once before each test, implement testify/suite's SetupTestSuite interface.
func (s *UserRepositorySuite) SetupTest() {
	s.T().Log("UserRepositorySuite SetupTest")
}

// TearDownTest run once before each test, implement testify/suite's TearDownTestSuite interface.
func (s *UserRepositorySuite) TearDownTest() {
	s.T().Log("UserRepositorySuite TearDownTest")

	sess := s.TestDB.Session(&gorm.Session{AllowGlobalUpdate: true})
	if err := sess.Unscoped().Delete(&domain.User{}).Error; err != nil {
		s.T().Fatal(err.Error())
	}
}
