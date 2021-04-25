package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/timwmillard/fishing/postgres/db"
)

type CompetitorQueries struct {
	// List
	ListCompetitorsFunc    func(ctx context.Context) ([]db.FishingCompetitor, error)
	ListCompetitorsInvoked bool

	// Get
	GetCompetitorFunc    func(ctx context.Context, id uuid.UUID) (db.FishingCompetitor, error)
	GetCompetitorInvoked bool

	// Create
	CreateCompetitorFunc    func(ctx context.Context, arg db.CreateCompetitorParams) (db.FishingCompetitor, error)
	CreateCompetitorInvoked bool

	// Update
	UpdateCompetitorFunc    func(ctx context.Context, arg db.UpdateCompetitorParams) (db.FishingCompetitor, error)
	UpdateCompetitorInvoked bool

	// Delete
	DeleteCompetitorFunc    func(ctx context.Context, id uuid.UUID) (int64, error)
	DeleteCompetitorInvoked bool
}

func (q *CompetitorQueries) CreateCompetitor(ctx context.Context, arg db.CreateCompetitorParams) (db.FishingCompetitor, error) {
	q.CreateCompetitorInvoked = true
	return q.CreateCompetitorFunc(ctx, arg)

}

func (q *CompetitorQueries) DeleteCompetitor(ctx context.Context, id uuid.UUID) (int64, error) {
	q.DeleteCompetitorInvoked = true
	return q.DeleteCompetitorFunc(ctx, id)
}

func (q *CompetitorQueries) GetCompetitor(ctx context.Context, id uuid.UUID) (db.FishingCompetitor, error) {
	q.GetCompetitorInvoked = true
	return q.GetCompetitorFunc(ctx, id)
}

func (q *CompetitorQueries) ListCompetitors(ctx context.Context) ([]db.FishingCompetitor, error) {
	q.ListCompetitorsInvoked = true
	return q.ListCompetitorsFunc(ctx)
}

func (q *CompetitorQueries) UpdateCompetitor(ctx context.Context, arg db.UpdateCompetitorParams) (db.FishingCompetitor, error) {
	q.UpdateCompetitorInvoked = true
	return q.UpdateCompetitorFunc(ctx, arg)
}