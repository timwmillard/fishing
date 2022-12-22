package fishing

import (
	"context"
	"errors"
)

// Competitor is a competitor in fishing competition.
type Competitor struct {
	ID           HashID `json:"id,omitempty"`
	CompetitorNo string `json:"competitor_no,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email,omitempty"`
	Address1     string `json:"address1,omitempty"`
	Address2     string `json:"address2,omitempty" `
	Suburb       string `json:"suburb,omitempty" `
	State        string `json:"state,omitempty" `
	Postcode     string `json:"postcode,omitempty" `
	Mobile       string `json:"mobile,omitempty" `
}

// Competitor is a competitor in fishing competition.
type CreateCompetitorParams struct {
	CompetitorNo string `json:"competitor_no,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email,omitempty"`
	Address1     string `json:"address1,omitempty"`
	Address2     string `json:"address2,omitempty"`
	Suburb       string `json:"suburb,omitempty"`
	State        string `json:"state,omitempty"`
	Postcode     string `json:"postcode,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
}

// Competitor is a competitor in fishing competition.
type UpdateCompetitorParams struct {
	CompetitorNo *string `json:"competitor_no,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
	Email        *string `json:"email,omitempty"`
	Address1     *string `json:"address1,omitempty"`
	Address2     *string `json:"address2,omitempty"`
	Suburb       *string `json:"suburb,omitempty"`
	State        *string `json:"state,omitempty"`
	Postcode     *string `json:"postcode,omitempty"`
	Mobile       *string `json:"mobile,omitempty"`
}

// CompetitorsRepo interface for competitors repository.
type CompetitorRepo interface {
	List(ctx context.Context) ([]Competitor, error)
	Get(ctx context.Context, id HashID) (Competitor, error)
	Create(ctx context.Context, c CreateCompetitorParams) (Competitor, error)
	Update(ctx context.Context, id HashID, c CreateCompetitorParams) (Competitor, error)
	UpdatePartial(ctx context.Context, id HashID, c UpdateCompetitorParams) (Competitor, error)
	Delete(ctx context.Context, id HashID) error
}

// Common Errors
var (
	ErrCompetitorNotFound = errors.New("competitor not found")
)
