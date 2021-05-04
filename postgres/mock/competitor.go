package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/timwmillard/fishing/postgres/sqlc"
)

var _ sqlc.Querier = (*CompetitorQueries)(nil)

type CompetitorQueries struct {
	// List
	ListCompetitorsFunc    func(ctx context.Context) ([]sqlc.FishingCompetitor, error)
	ListCompetitorsInvoked bool

	// Get
	GetCompetitorFunc    func(ctx context.Context, id uuid.UUID) (sqlc.FishingCompetitor, error)
	GetCompetitorInvoked bool

	// Create
	CreateCompetitorFunc    func(ctx context.Context, arg sqlc.CreateCompetitorParams) (sqlc.FishingCompetitor, error)
	CreateCompetitorInvoked bool

	// Update
	UpdateCompetitorFunc    func(ctx context.Context, arg sqlc.UpdateCompetitorParams) (sqlc.FishingCompetitor, error)
	UpdateCompetitorInvoked bool

	// Delete
	DeleteCompetitorFunc    func(ctx context.Context, id uuid.UUID) (int64, error)
	DeleteCompetitorInvoked bool
}

func (q *CompetitorQueries) CreateCompetitor(ctx context.Context, arg sqlc.CreateCompetitorParams) (sqlc.FishingCompetitor, error) {
	q.CreateCompetitorInvoked = true
	return q.CreateCompetitorFunc(ctx, arg)

}

func (q *CompetitorQueries) DeleteCompetitor(ctx context.Context, id uuid.UUID) (int64, error) {
	q.DeleteCompetitorInvoked = true
	return q.DeleteCompetitorFunc(ctx, id)
}

func (q *CompetitorQueries) GetCompetitor(ctx context.Context, id uuid.UUID) (sqlc.FishingCompetitor, error) {
	q.GetCompetitorInvoked = true
	return q.GetCompetitorFunc(ctx, id)
}

func (q *CompetitorQueries) ListCompetitors(ctx context.Context) ([]sqlc.FishingCompetitor, error) {
	q.ListCompetitorsInvoked = true
	return q.ListCompetitorsFunc(ctx)
}

func (q *CompetitorQueries) UpdateCompetitor(ctx context.Context, arg sqlc.UpdateCompetitorParams) (sqlc.FishingCompetitor, error) {
	q.UpdateCompetitorInvoked = true
	return q.UpdateCompetitorFunc(ctx, arg)
}

func Competitor() sqlc.FishingCompetitor {
	return sqlc.FishingCompetitor{
		ID:        uuid.New(),
		Firstname: fake.FirstName(),
		Lastname:  fake.LastName(),
		Email:     fake.EmailAddress(),
		Address1:  fake.StreetAddress(),
		Suburb:    fake.City(),
		Postcode:  fake.DigitsN(4),
		Phone:     fake.Phone(),
		Mobile:    fake.Phone(),
	}
}

func Competitors(n int) []sqlc.FishingCompetitor {
	comps := make([]sqlc.FishingCompetitor, n)
	for i := 0; i < n; i++ {
		comps[i] = Competitor()
	}
	return comps
}
