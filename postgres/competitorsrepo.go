// Postgres implementation of CompetitorsRepo
//
//go:generate sqlc generate

package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fishing"
	"fishing/postgres/db"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql" // Need MySQL driver
	"github.com/jmoiron/sqlx"
)

// CompetitorsRepo -
type CompetitorsRepo struct {
	db    *sqlx.DB
	query *db.Queries
}

// NewCompetitorsRepo -
func NewCompetitorsRepo(connection *sqlx.DB) *CompetitorsRepo {

	return &CompetitorsRepo{
		db:    connection,
		query: db.New(connection),
	}
}

// List -
func (r *CompetitorsRepo) List() ([]*fishing.Competitor, error) {

	comp, err := r.query.ListCompetitors(context.TODO())
	if err != nil {
		return nil, err
	}

	fishComp := fishingCompetitors(comp)

	return fishComp, nil
}

// Get -
func (r *CompetitorsRepo) Get(id uuid.UUID) (*fishing.Competitor, error) {
	var competitor fishing.Competitor

	q := `SELECT uuid, event_id, competitor_no, firstname, lastname, email, address1, address2, suburb, state, postcode, phone, mobile
		  FROM competitors WHERE uuid=UUID_TO_BIN(?)`
	err := r.db.Get(&competitor, q, id)
	if err != nil {
		return nil, err
	}

	return &competitor, nil
}

// Create -
func (r *CompetitorsRepo) Create(c *fishing.Competitor) (*fishing.Competitor, error) {
	q := `INSERT INTO competitors
		  		(uuid, event_id, competitor_no, firstname, lastname, email, address1, address2, suburb, state, postcode, phone, mobile)
		  VALUE (UUID_TO_BIN(UUID()), :event_id, :competitor_no, :firstname, :lastname, :email, :address1, :address2, :suburb, :state, :postcode, :phone, :mobile)`
	result, err := r.db.NamedExec(q, c)
	if err != nil {
		return nil, err
	}

	// Update the competitor wit the new ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var newCompetitor fishing.Competitor
	q = `SELECT uuid, event_id, competitor_no, firstname, lastname, email, address1, address2, suburb, state, postcode, phone, mobile
		  FROM competitors WHERE id=?`
	err = r.db.Get(&newCompetitor, q, id)
	if err != nil {
		return nil, err
	}

	return &newCompetitor, nil
}

// Update -
func (r *CompetitorsRepo) Update(id uuid.UUID, c *fishing.Competitor) (*fishing.Competitor, error) {
	c.ID = id

	q := `UPDATE competitors
		  SET event_id=:event_id, competitor_no=:competitor_no, firstname=:firstname, lastname=:lastname, email=:email, address1=:address1, address2=:address2,
		  	 suburb=:suburb, state=:state, postcode=:postcode, phone=:phone, mobile=:mobile
		  WHERE uuid = UUID_TO_BIN(:uuid)`
	result, err := r.db.NamedExec(q, c)
	if err != nil {
		return nil, err
	}

	numUpdated, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if numUpdated < 1 {
		return nil, errors.New("competitor not found")
	}

	return c, nil
}

// Delete -
func (r *CompetitorsRepo) Delete(id uuid.UUID) error {
	q := `DELETE FROM competitors
		  WHERE uuid = UUID_TO_BIN(?)`
	result, err := r.db.Exec(q, id)
	if err != nil {
		return err
	}

	numDeleted, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if numDeleted < 1 {
		return errors.New("competitor not found")
	}

	return nil
}

func fishingCompetitors(dbComps []db.Competitor) []*fishing.Competitor {
	fishComps := make([]*fishing.Competitor, len(dbComps))
	for _, c := range dbComps {
		p := fishingCompetitor(c)
		fishComps = append(fishComps, &p)
	}
	return fishComps
}

func fishingCompetitor(c db.Competitor) fishing.Competitor {
	return fishing.Competitor{
		ID:           c.ID,
		CompetitorNo: nullInt(c.CompetitorNo),
		Firstname:    c.Firstname,
		Lastname:     c.Lastname,
		Email:        c.Email,
		Address1:     c.Address1,
		Address2:     c.Address2,
		Suburb:       c.Suburb,
		State:        c.State,
		Postcode:     c.Postcode,
		Phone:        c.Phone,
		Mobile:       c.Mobile,
		EventID:      nullInt(c.EventID),
	}
}

func nullInt(i sql.NullInt32) *int {
	_, err := i.Value()
	if err != nil {
		return nil
	}
	r := int(i.Int32)
	return &r
}

//lint:ignore U1000 unused utility function
// TODO write test
func nullString(s sql.NullString) *string {
	_, err := s.Value()
	if err != nil {
		return nil
	}
	r := string(s.String)
	return &r
}
