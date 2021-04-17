// Postgres implementation of CompetitorsRepo
package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/timwmillard/fishing"
	"github.com/timwmillard/fishing/postgres/db"
)

// CompetitorsRepo -
type CompetitorsRepo struct {
	query *db.Queries
}

// NewCompetitorsRepo -
func NewCompetitorsRepo(connection *sql.DB) *CompetitorsRepo {
	return &CompetitorsRepo{
		query: db.New(connection),
	}
}

// List -
func (r *CompetitorsRepo) List(ctx context.Context) ([]fishing.Competitor, error) {
	comps, err := r.query.ListCompetitors(context.TODO())
	if err != nil {
		return nil, err
	}
	return competitors(comps), nil

}

// Get -
func (r *CompetitorsRepo) Get(ctx context.Context, id uuid.UUID) (fishing.Competitor, error) {
	comp, err := r.query.GetCompetitor(ctx, id)
	if err != nil {
		return fishing.Competitor{}, err
	}
	return competitor(comp), nil
}

// Create -
func (r *CompetitorsRepo) Create(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	comp, err := r.query.CreateCompetitor(ctx, createCompetitorParams(c))
	if err != nil {
		return fishing.Competitor{}, nil
	}
	return competitor(comp), nil
}

// Update -
func (r *CompetitorsRepo) Update(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	comp, err := r.query.UpdateCompetitor(ctx, updateCompetitorParams(c))
	if err != nil {
		return fishing.Competitor{}, nil
	}
	return competitor(comp), nil
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
		ID:           c.ID,
		CompetitorNo: c.CompetitorNo,
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
	}
}

func updateCompetitorParams(c fishing.Competitor) db.UpdateCompetitorParams {
	return db.UpdateCompetitorParams{
		ID:           c.ID,
		CompetitorNo: c.CompetitorNo,
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
	}
}

func competitors(dbComps []db.FishingCompetitor) []fishing.Competitor {
	fishComps := make([]fishing.Competitor, len(dbComps))
	for _, c := range dbComps {
		p := competitor(c)
		fishComps = append(fishComps, p)
	}
	return fishComps
}

func competitor(c db.FishingCompetitor) fishing.Competitor {
	return fishing.Competitor{
		ID:           c.ID,
		CompetitorNo: c.CompetitorNo,
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
		// EventID:      c.EventID,
	}
}
