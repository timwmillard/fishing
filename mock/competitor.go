package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/timwmillard/fishing"
)

type CompetitorMock struct {

	// List
	ListCtxIn          context.Context
	ListCompetitorsOut []fishing.Competitor
	ListErrorOut       error

	// Get
	GetFunc    func(context.Context, uuid.UUID) (fishing.Competitor, error)
	GetInvoked bool

	// Create
	CreateFunc    func(context.Context, fishing.Competitor) (fishing.Competitor, error)
	CreateInvoked bool

	// Update
	UpdateFunc    func(context.Context, fishing.Competitor) (fishing.Competitor, error)
	UpdateInvoked bool

	// Delete
	DeleteFunc    func(context.Context, uuid.UUID) error
	DeleteInvoked bool
}

func (cm *CompetitorMock) List(ctx context.Context) ([]fishing.Competitor, error) {
	cm.ListCtxIn = ctx

	return cm.ListCompetitorsOut, cm.ListErrorOut
}

func (cm *CompetitorMock) Get(ctx context.Context, id uuid.UUID) (fishing.Competitor, error) {
	cm.GetInvoked = true
	return cm.GetFunc(ctx, id)
}

func (cm *CompetitorMock) Create(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	cm.CreateInvoked = true
	return cm.CreateFunc(ctx, c)
}

func (cm *CompetitorMock) Update(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	cm.UpdateInvoked = true
	return cm.UpdateFunc(ctx, c)
}

func (cm *CompetitorMock) Delete(ctx context.Context, id uuid.UUID) error {
	cm.DeleteInvoked = true
	return cm.DeleteFunc(ctx, id)
}
