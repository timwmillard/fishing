package testfixture

import (
	"database/sql"

	"github.com/go-testfixtures/testfixtures"
)

func New() {
	var err error

	// Open connection to the test database.
	// Do NOT import fixtures in a production database!
	// Existing data would be deleted.
	db, err = sql.Open("postgres", "dbname=myapp_test")
	if err != nil {

	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),                   // You database connection
		testfixtures.Dialect("postgres"),            // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("testdata/fixtures"), // the directory containing the YAML files
	)
	if err != nil {

	}
}
