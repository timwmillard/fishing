package postgres_test

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/assert"
	"github.com/timwmillard/fishing/fake"
	"github.com/timwmillard/fishing/postgres"
)

var (
	db       *sql.DB
	database = "fishingcomp"

	competitorsRepo *postgres.CompetitorsRepo
)

var comp1 = fake.Competitor()

func TestMain(m *testing.M) {
	flag.Parse()
	if testing.Short() {
		m.Run()
		os.Exit(0)
	}
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12-alpine",
		Env: []string{
			"POSTGRES_USER=root",
			"POSTGRES_PASSWORD=fish",
			"POSTGRES_DB=" + database,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: "5999"},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", fmt.Sprintf("postgres://root:fish@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), database))
		if err != nil {
			return err
		}

		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	driver, err := migratepg.WithInstance(db, &migratepg.Config{})
	if err != nil {
		log.Fatalf("Could not migrate driver: %s", err)
	}
	mig, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Could not migrate new database instance: %s", err)
	}
	mig.Up()

	competitorsRepo = postgres.NewCompetitorsRepo(db)

	code := m.Run()

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}

func TestCreate(t *testing.T) {
	// db.Query()
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	ctx := context.Background()

	c1, err := competitorsRepo.Create(ctx, comp1)
	assert.NoError(t, err)
	assert.Equal(t, c1.Firstname, comp1.Firstname)
	assert.Equal(t, c1.Lastname, comp1.Lastname)
	assert.Equal(t, c1.Email, comp1.Email)

	c2, err := competitorsRepo.Get(ctx, c1.ID)
	assert.NoError(t, err)
	assert.Equal(t, c2.Firstname, comp1.Firstname)
	assert.Equal(t, c2.Lastname, comp1.Lastname)
}
