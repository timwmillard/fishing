package fishing

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

// Competitor is a competitor in fishing competition.
type Competitor struct {
	ID           uuid.UUID `json:"id,omitempty" db:"id"`
	CompetitorNo string    `json:"competitor_no,omitempty" db:"competitor_no"`
	Firstname    string    `json:"first_name,omitempty" db:"firstname"`
	Lastname     string    `json:"last_name,omitempty" db:"lastname"`
	Email        string    `json:"email,omitempty" db:"email"`
	Address1     string    `json:"address1,omitempty" db:"address1"`
	Address2     string    `json:"address2,omitempty"  db:"address2"`
	Suburb       string    `json:"suburb,omitempty"  db:"suburb"`
	State        string    `json:"state,omitempty"  db:"state"`
	Postcode     string    `json:"postcode,omitempty"  db:"postcode"`
	Phone        string    `json:"phone,omitempty"  db:"phone"`
	Mobile       string    `json:"mobile,omitempty"  db:"mobile"`
	EventID      uuid.UUID `json:"event_id,omitempty" db:"event_id"`
}

// CompetitorsRepo interface for competitors repository.
type CompetitorsRepo interface {
	List(ctx context.Context) ([]Competitor, error)
	Get(ctx context.Context, id uuid.UUID) (Competitor, error)
	Create(ctx context.Context, c Competitor) (Competitor, error)
	Update(ctx context.Context, c Competitor) (Competitor, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// Common Errors
var (
	ErrCompetitorNotFound = errors.New("competitor not found")
)
