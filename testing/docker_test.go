// +build integration

package testing

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/matryer/is"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"

	"github.com/timwmillard/fishing/fake"
	"github.com/timwmillard/fishing/postgres"
)

const (
	dockerPostgresUser     = "root"
	dockerPostgresPassword = "fish"
	dockerDatabase         = "fishingcomp"
)

var (
	db             *sql.DB
	competitorRepo *postgres.CompetitorRepo
)

func TestMain(m *testing.M) {
	err := setupDockerPostres()
	if err != nil {
		log.Fatalf("setup docker postres: %v", err)
	}
	code := m.Run()
	os.Exit(code)
}

func setupDockerPostres() error {
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		return fmt.Errorf("could not connect to docker: %v", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12-alpine",
		Env: []string{
			"POSTGRES_USER=" + dockerPostgresUser,
			"POSTGRES_PASSWORD=" + dockerPostgresPassword,
			"POSTGRES_DB=" + dockerDatabase,
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
		return fmt.Errorf("could not start resource: %v", err)
	}
	defer func() {
		// When you're done, kill and remove the container
		if err = pool.Purge(resource); err != nil {
			log.Printf("could not purge resource: %v", err)
		}
	}()

	if err = pool.Retry(func() error {
		var err error
		source := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", dockerPostgresUser, dockerPostgresPassword, resource.GetPort("5432/tcp"), dockerDatabase)
		db, err = sql.Open("postgres", source)
		if err != nil {
			return err
		}

		return db.Ping()
	}); err != nil {
		return fmt.Errorf("connect to postgres: %v", err)
	}

	driver, err := migratepg.WithInstance(db, &migratepg.Config{})
	if err != nil {
		return fmt.Errorf("migrate driver: %v", err)
	}
	mig, err := migrate.NewWithDatabaseInstance(
		"file://../postgres/migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("migrate new database instance: %v", err)
	}
	mig.Up()

	competitorRepo = postgres.NewCompetitorRepo(db)

	return nil
}

func TestCompetitorRepo_Create(t *testing.T) {

	is := is.New(t)
	ctx := context.Background()

	comp1 := fake.Competitor()

	c1, err := competitorRepo.Create(ctx, comp1)
	is.NoErr(err)
	is.Equal(c1.Firstname, comp1.Firstname)
	is.Equal(c1.Lastname, comp1.Lastname)
	is.Equal(c1.Email, comp1.Email)
	is.True(reflect.DeepEqual(c1, comp1))

	c2, err := competitorRepo.Get(ctx, c1.ID)
	is.NoErr(err)
	is.True(reflect.DeepEqual(c2, comp1))
}

// func TestCompetitorList_docker(t *testing.T) {
// 	// db.Query()
// 	if testing.Short() {
// 		t.Skip("skipping test in short mode.")
// 	}

// 	is := is.New(t)
// 	ctx := context.Background()

// 	want := fake.Competitors(5)
// 	for _, comp := range want {
// 		_, err := competitorRepo.Create(ctx, comp)
// 		is.NoErr(err)
// 	}

// 	got, err := competitorRepo.List(ctx)
// 	is.NoErr(err)
// 	is.Equal(len(got), len(want))
// }
