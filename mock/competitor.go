package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/timwmillard/fishing"
)

var _ fishing.CompetitorRepo = (*CompetitorRepo)(nil)

type CompetitorRepo struct {

	// List
	ListFunc    func(ctx context.Context) ([]fishing.Competitor, error)
	ListInvoked bool

	// Get
	GetFunc    func(ctx context.Context, id uuid.UUID) (fishing.Competitor, error)
	GetInvoked bool

	// Create
	CreateFunc    func(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error)
	CreateInvoked bool

	// Update
	UpdateFunc    func(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error)
	UpdateInvoked bool

	// Delete
	DeleteFunc    func(ctx context.Context, id uuid.UUID) error
	DeleteInvoked bool
}

func (cm *CompetitorRepo) List(ctx context.Context) ([]fishing.Competitor, error) {
	cm.ListInvoked = true
	return cm.ListFunc(ctx)
}

func (cm *CompetitorRepo) Get(ctx context.Context, id uuid.UUID) (fishing.Competitor, error) {
	cm.GetInvoked = true
	return cm.GetFunc(ctx, id)
}

func (cm *CompetitorRepo) Create(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	cm.CreateInvoked = true
	return cm.CreateFunc(ctx, c)
}

func (cm *CompetitorRepo) Update(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	cm.UpdateInvoked = true
	return cm.UpdateFunc(ctx, c)
}

func (cm *CompetitorRepo) Delete(ctx context.Context, id uuid.UUID) error {
	cm.DeleteInvoked = true
	return cm.DeleteFunc(ctx, id)
}

func Competitor() fishing.Competitor {
	return fishing.Competitor{
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

func Competitors(n int) []fishing.Competitor {
	comps := make([]fishing.Competitor, n)
	for i := 0; i < n; i++ {
		comps[i] = Competitor()
	}
	return comps
}
