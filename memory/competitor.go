package memory

import (
	"context"

	"github.com/google/uuid"
	"github.com/timwmillard/fishing"
)

var _ fishing.CompetitorRepo = (*CompetitorRepo)(nil)

// CompetitorsRepo -
type CompetitorRepo struct {
	data map[uuid.UUID]fishing.Competitor
}

// NewCompetitorRepo -
func NewCompetitorRepo() *CompetitorRepo {
	data := make(map[uuid.UUID]fishing.Competitor)
	repo := &CompetitorRepo{data: data}
	return repo
}

// List -
func (r *CompetitorRepo) List(ctx context.Context) ([]fishing.Competitor, error) {
	var list []fishing.Competitor
	for _, competitor := range r.data {
		list = append(list, competitor)
	}
	return list, nil
}

// Get -
func (r *CompetitorRepo) Get(ctx context.Context, id uuid.UUID) (fishing.Competitor, error) {
	competitor, ok := r.data[id]
	if !ok {
		return fishing.Competitor{}, fishing.ErrCompetitorNotFound
	}
	return competitor, nil
}

// Create -
func (r *CompetitorRepo) Create(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	id := uuid.New()
	c.ID = id
	r.data[id] = c
	return c, nil
}

// Update -
func (r *CompetitorRepo) Update(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	_, ok := r.data[c.ID]
	if !ok {
		return fishing.Competitor{}, fishing.ErrCompetitorNotFound
	}

	r.data[c.ID] = c
	return c, nil
}

// Delete -
func (r *CompetitorRepo) Delete(ctx context.Context, id uuid.UUID) error {
	delete(r.data, id)
	return nil
}
