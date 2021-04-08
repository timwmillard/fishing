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
	comp, err := r.query.CreateCompetitor(ctx, createCompetitorParams(c))
	if err != nil {
		return fishing.Competitor{}, nil
	}
	return fishingCompetitor(comp), nil
}

// Update -
func (r *CompetitorsRepo) Update(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	comp, err := r.query.UpdateCompetitor(ctx, updateCompetitorParams(c))
	if err != nil {
		return fishing.Competitor{}, nil
	}
	return fishingCompetitor(comp), nil
}

// Delete -
func (r *CompetitorsRepo) Delete(ctx context.Context, id uuid.UUID) error {

	numDeleted, err := r.query.DeleteCompetitor(ctx, id)
	if err != nil {
		return err
	}
	if numDeleted < 1 {
		return fishing.ErrCompetitorNotFound
	}

	return nil
}

func createCompetitorParams(c fishing.Competitor) db.CreateCompetitorParams {
	return db.CreateCompetitorParams{
		ID:        c.ID,
		Firstname: c.Firstname,
		Lastname:  c.Lastname,
	}
}

func updateCompetitorParams(c fishing.Competitor) db.UpdateCompetitorParams {
	return db.UpdateCompetitorParams{
		ID:        c.ID,
		Firstname: c.Firstname,
		Lastname:  c.Lastname,
	}
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
