package dockertest

import (
	"database/sql"
	"fmt"
	"strconv"

	// Required for sql.Open postgres.
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"gitlab.com/trivery-id/skadi/external/db/postgres"
	"gitlab.com/trivery-id/skadi/utils/errors"
)

type PostgreSQLPool struct {
	pool     *dockertest.Pool
	resource *dockertest.Resource
	config   *setupOptions

	Credential postgres.DBCredential
}

func NewPostgreSQLPool(opts ...SetupOption) (*PostgreSQLPool, error) {
	options := parseOptions(opts...)

	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, errors.Newf("dockertest: could not connect to docker: %v", err)
	}

	resource, err := pool.Run(
		options.DockerImage,
		options.DockerImageTag,
		[]string{
			fmt.Sprintf("POSTGRES_PASSWORD=%s", options.DBPassword),
			fmt.Sprintf("POSTGRES_DB=%s", options.DBName),
		},
	)
	if err != nil {
		return nil, errors.Newf("dockertest: could not start resource: %v", err)
	}

	dbPort, err := strconv.Atoi(resource.GetPort("5432/tcp"))
	if err != nil {
		return nil, errors.New("dockertest: could not connect to docker: invalid port")
	}

	pgConnString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		options.DockerHost,
		dbPort,
		options.DBUsername,
		options.DBPassword,
	)
	if err := pool.Retry(func() error {
		db, err := sql.Open("postgres", pgConnString)
		if err != nil {
			return err
		}

		return db.Ping()
	}); err != nil {
		return nil, errors.Newf("dockertest: could not connect to docker: %v", err)
	}

	return &PostgreSQLPool{
		pool:     pool,
		resource: resource,
		config:   options,
		Credential: postgres.DBCredential{
			Host:     options.DockerHost,
			Port:     dbPort,
			Username: options.DBUsername,
			Password: options.DBPassword,
			DBName:   options.DBName,
		},
	}, nil
}

// Purge kill and remove the container.
func (p *PostgreSQLPool) Purge() error {
	if err := p.pool.Purge(p.resource); err != nil {
		return errors.Newf("dockertest: could not purge resource: %v", err)
	}

	return nil
}
