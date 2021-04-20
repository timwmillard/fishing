package postgres

import (
	"context"
	"database/sql"
	"log"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/matryer/is"
	"github.com/timwmillard/fishing/fake"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

const getCompetitor = `-- name: GetCompetitor :one
SELECT id, competitor_no, firstname, lastname, email, address1, address2, suburb, state, postcode, phone, mobile FROM competitors
WHERE id = $1
`

func TestGet(t *testing.T) {
	is := is.New(t)

	comp1 := fake.Competitor()

	db, mock := NewMock()
	rows := sqlmock.NewRows([]string{"id", "competitor_no", "firstname", "lastname", "email", "address1", "address2", "suburb", "state", "postcode", "phone", "mobile"}).
		AddRow(
			comp1.ID,
			comp1.CompetitorNo,
			comp1.Firstname,
			comp1.Lastname,
			comp1.Email,
			comp1.Address1,
			comp1.Address2,
			comp1.Suburb,
			comp1.State,
			comp1.Postcode,
			comp1.Phone,
			comp1.Mobile,
		)

	result := mock.ExpectQuery(regexp.QuoteMeta(getCompetitor)).WithArgs(comp1.ID).WillReturnRows(rows)
	t.Log("Expected Query: ", result)

	ctx := context.Background()

	repo := NewCompetitorRepo(db)
	defer func() {
		db.Close()
	}()

	got, err := repo.Get(ctx, comp1.ID)
	is.NoErr(err)
	is.True(reflect.DeepEqual(got, comp1))

}
