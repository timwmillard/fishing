// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateCompetitor(ctx context.Context, arg CreateCompetitorParams) (FishingCompetitor, error)
	DeleteCompetitor(ctx context.Context, id int64) (int64, error)
	GetCompetitor(ctx context.Context, id int64) (FishingCompetitor, error)
	ListCompetitor(ctx context.Context) ([]FishingCompetitor, error)
	UpdateCompetitor(ctx context.Context, arg UpdateCompetitorParams) (FishingCompetitor, error)
}

var _ Querier = (*Queries)(nil)
