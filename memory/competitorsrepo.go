package memory

import (
	"errors"
	"fishing"

	"github.com/google/uuid"
)

// CompetitorsRepo -
type CompetitorsRepo struct {
	data map[uuid.UUID]*fishing.Competitor
}

// NewCompetitorsRepo -
func NewCompetitorsRepo() *CompetitorsRepo {
	data := make(map[uuid.UUID]*fishing.Competitor)
	repo := &CompetitorsRepo{data: data}
	return repo
}

// List -
func (r *CompetitorsRepo) List() ([]*fishing.Competitor, error) {
	var list []*fishing.Competitor
	for _, competitor := range r.data {
		list = append(list, competitor)
	}
	return list, nil
}

// Get -
func (r *CompetitorsRepo) Get(id uuid.UUID) (*fishing.Competitor, error) {
	competitor, ok := r.data[id]
	if !ok {
		return nil, errors.New("competitor not found")
	}
	return competitor, nil
}

// Create -
func (r *CompetitorsRepo) Create(c *fishing.Competitor) (*fishing.Competitor, error) {
	id := uuid.New()
	c.ID = id
	r.data[id] = c
	return c, nil
}

// Update -
func (r *CompetitorsRepo) Update(id uuid.UUID, c *fishing.Competitor) (*fishing.Competitor, error) {
	_, ok := r.data[id]
	if !ok {
		return nil, errors.New("competitor not found")
	}
	c.ID = id
	r.data[id] = c
	return c, nil
}

// Delete -
func (r *CompetitorsRepo) Delete(id uuid.UUID) error {
	delete(r.data, id)
	return nil
}
