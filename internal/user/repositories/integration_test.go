package repositories_test

import (
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"gitlab.com/trivery-id/skadi/external/db/postgres"
	"gitlab.com/trivery-id/skadi/utils/dockertest"
)

var cred postgres.DBCredential

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPostgreSQLPool()
	if err != nil {
		log.Fatal(err.Error())
	}

	cred = pool.Credential
	exitCode := m.Run()

	if err := pool.Purge(); err != nil {
		log.Fatal(err.Error())
	}

	os.Exit(exitCode)
}
