package memory

import (
	"context"
	"fishing"

	"github.com/google/uuid"
)

// CompetitorsRepo -
type CompetitorsRepo struct {
	data map[uuid.UUID]fishing.Competitor
}

// NewCompetitorsRepo -
func NewCompetitorsRepo() *CompetitorsRepo {
	data := make(map[uuid.UUID]fishing.Competitor)
	repo := &CompetitorsRepo{data: data}
	return repo
}

// List -
func (r *CompetitorsRepo) List(ctx context.Context) ([]fishing.Competitor, error) {
	var list []fishing.Competitor
	for _, competitor := range r.data {
		list = append(list, competitor)
	}
	return list, nil
}

// Get -
func (r *CompetitorsRepo) Get(ctx context.Context, id uuid.UUID) (fishing.Competitor, error) {
	competitor, ok := r.data[id]
	if !ok {
		return fishing.Competitor{}, fishing.ErrCompetitorNotFound
	}
	return competitor, nil
}

// Create -
func (r *CompetitorsRepo) Create(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	id := uuid.New()
	c.ID = id
	r.data[id] = c
	return c, nil
}

// Update -
func (r *CompetitorsRepo) Update(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	_, ok := r.data[c.ID]
	if !ok {
		return fishing.Competitor{}, fishing.ErrCompetitorNotFound
	}

	r.data[c.ID] = c
	return c, nil
}

// Delete -
func (r *CompetitorsRepo) Delete(ctx context.Context, id uuid.UUID) error {
	delete(r.data, id)
	return nil
}
