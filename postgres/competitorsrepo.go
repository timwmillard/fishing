// Postgres implementation of CompetitorsRepo
//
//go:generate sqlc generate

package postgres

import (
	"context"
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
func (r *CompetitorsRepo) List(ctx context.Context) ([]fishing.Competitor, error) {
	comps, err := r.query.ListCompetitors(context.TODO())
	if err != nil {
		return nil, err
	}
	return fishingCompetitors(comps), nil

}

// Get -
func (r *CompetitorsRepo) Get(ctx context.Context, id uuid.UUID) (fishing.Competitor, error) {
	comp, err := r.query.GetCompetitor(ctx, id)
	if err != nil {
		return fishing.Competitor{}, err
	}
	return fishingCompetitor(comp), nil
}

// Create -
func (r *CompetitorsRepo) Create(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	q := `INSERT INTO competitors
		  		(uuid, event_id, competitor_no, firstname, lastname, email, address1, address2, suburb, state, postcode, phone, mobile)
		  VALUE (UUID_TO_BIN(UUID()), :event_id, :competitor_no, :firstname, :lastname, :email, :address1, :address2, :suburb, :state, :postcode, :phone, :mobile)`
	result, err := r.db.NamedExec(q, c)
	if err != nil {
		return fishing.Competitor{}, err
	}

	// Update the competitor wit the new ID
	id, err := result.LastInsertId()
	if err != nil {
		return fishing.Competitor{}, err
	}

	var newCompetitor fishing.Competitor
	q = `SELECT uuid, event_id, competitor_no, firstname, lastname, email, address1, address2, suburb, state, postcode, phone, mobile
		  FROM competitors WHERE id=?`
	err = r.db.Get(&newCompetitor, q, id)
	if err != nil {
		return fishing.Competitor{}, err
	}

	return newCompetitor, nil
}

// Update -
func (r *CompetitorsRepo) Update(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {

	q := `UPDATE competitors
		  SET event_id=:event_id, competitor_no=:competitor_no, firstname=:firstname, lastname=:lastname, email=:email, address1=:address1, address2=:address2,
		  	 suburb=:suburb, state=:state, postcode=:postcode, phone=:phone, mobile=:mobile
		  WHERE uuid = UUID_TO_BIN(:uuid)`
	result, err := r.db.NamedExec(q, c)
	if err != nil {
		return fishing.Competitor{}, err
	}

	numUpdated, err := result.RowsAffected()
	if err != nil {
		return fishing.Competitor{}, err
	}
	if numUpdated < 1 {
		return fishing.Competitor{}, fishing.ErrCompetitorNotFound
	}

	return c, nil
}

// Delete -
func (r *CompetitorsRepo) Delete(ctx context.Context, id uuid.UUID) error {
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
		return fishing.ErrCompetitorNotFound
	}

	return nil
}

func fishingCompetitors(dbComps []db.Competitor) []fishing.Competitor {
	fishComps := make([]fishing.Competitor, len(dbComps))
	for _, c := range dbComps {
		p := fishingCompetitor(c)
		fishComps = append(fishComps, p)
	}
	return fishComps
}

func fishingCompetitor(c db.Competitor) fishing.Competitor {
	return fishing.Competitor{
		ID:           c.ID,
		CompetitorNo: nullString(c.CompetitorNo),
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
		EventID:      c.EventID,
	}
}
