// Postgres implementation of CompetitorRepo.
package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/timwmillard/fishing"
	"github.com/timwmillard/fishing/postgres/db"
)

// CompetitorRepo is a repository of competitors.
type CompetitorRepo struct {
	query db.Querier
}

// NewCompetitorsRepo returns a new competitor repository.
// connection is a existing sql.DB connection.
func NewCompetitorRepo(connection *sql.DB) *CompetitorRepo {
	return &CompetitorRepo{
		query: db.New(connection),
	}
}

// newCompetitorRepoWithQuerier used to add a mock Querier for testing.
func newCompetitorRepoWithQuerier(q db.Querier) *CompetitorRepo {
	return &CompetitorRepo{
		query: q,
	}
}

// List returns a list of all competitors.
func (r *CompetitorRepo) List(ctx context.Context) ([]fishing.Competitor, error) {
	comps, err := r.query.ListCompetitors(ctx)
	if err != nil {
		return nil, err
	}
	return competitors(comps), nil

}

// Get's a single competitor by id.
func (r *CompetitorRepo) Get(ctx context.Context, id uuid.UUID) (fishing.Competitor, error) {
	comp, err := r.query.GetCompetitor(ctx, id)
	if err != nil {
		return fishing.Competitor{}, err
	}
	return competitor(comp), nil
}

// Create's a new competitor.
func (r *CompetitorRepo) Create(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	comp, err := r.query.CreateCompetitor(ctx, createCompetitorParams(c))
	if err != nil {
		return fishing.Competitor{}, err
	}
	return competitor(comp), nil
}

// Update's an existing competitor.  Returns the updated competitor.
func (r *CompetitorRepo) Update(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	comp, err := r.query.UpdateCompetitor(ctx, updateCompetitorParams(c))
	if err != nil {
		return fishing.Competitor{}, err
	}
	return competitor(comp), nil
}

// Delete's a competitor by id.
func (r *CompetitorRepo) Delete(ctx context.Context, id uuid.UUID) error {

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
	return db.CreateCompetitorParams(c)
}

func updateCompetitorParams(c fishing.Competitor) db.UpdateCompetitorParams {
	return db.UpdateCompetitorParams(c)
}

func competitors(dbComps []db.FishingCompetitor) []fishing.Competitor {
	fishComps := make([]fishing.Competitor, 0, len(dbComps))
	for _, c := range dbComps {
		p := competitor(c)
		fishComps = append(fishComps, p)
	}
	return fishComps
}

func competitor(c db.FishingCompetitor) fishing.Competitor {
	return fishing.Competitor(c)
}
