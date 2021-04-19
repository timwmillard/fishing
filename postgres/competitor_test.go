package postgres

import (
	"context"
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/timwmillard/fishing"
)

var comp1 = fishing.Competitor{
	ID:           uuid.Must(uuid.NewRandom()),
	CompetitorNo: "12",
	Firstname:    "Tim",
	Lastname:     "Millard",
	Email:        "tim@example.com",
	Address1:     "123 Main St",
	Address2:     "",
	Suburb:       "Some Town",
	State:        "VIC",
	Postcode:     "3000",
	Phone:        "123456",
	Mobile:       "04123456",
}

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

	mock.ExpectQuery(regexp.QuoteMeta(getCompetitor)).WithArgs(comp1.ID).WillReturnRows(rows)

	ctx := context.Background()

	repo := NewCompetitorsRepo(db)
	defer func() {
		db.Close()
	}()

	got, err := repo.Get(ctx, comp1.ID)
	assert.NotNil(t, got)
	assert.NoError(t, err)
}
