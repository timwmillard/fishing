package cache

import (
	"context"

	"github.com/google/uuid"
	"github.com/timwmillard/fishing"
)

type Competitors struct {
	Cache fishing.CompetitorRepo
	Repo  fishing.CompetitorRepo
}

func (cc *Competitors) List(ctx context.Context) ([]fishing.Competitor, error) {
	return cc.Repo.List(ctx)
}

func (cc *Competitors) Get(ctx context.Context, id uuid.UUID) (fishing.Competitor, error) {
	comp, err := cc.Cache.Get(ctx, id)
	if err == nil {
		return comp, nil
	}

	comp, err = cc.Repo.Get(ctx, id)
	if err != nil {
		return fishing.Competitor{}, err
	}
	cc.Cache.Create(ctx, comp)

	return comp, nil
}

func (cc *Competitors) Create(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	return cc.Repo.Create(ctx, c)
}

func (cc *Competitors) Update(ctx context.Context, c fishing.Competitor) (fishing.Competitor, error) {
	return cc.Repo.Update(ctx, c)
}

func (cc *Competitors) Delete(ctx context.Context, id uuid.UUID) error {
	return cc.Repo.Delete(ctx, id)
}
