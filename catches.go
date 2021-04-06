package fishing

import (
	"time"

	"github.com/google/uuid"
)

// Catch -
type Catch struct {
	UUID         uuid.UUID
	EventID      int
	CompetitorID int
	Species      Species
	Size         int // size in mm
	CaughtAt     time.Time
	Bait         string
	Location     string
	Latitude     float64
	Longitude    float64
	Marshall     string
	MarshallID   int
	Status       int
}

// CatchesRepo interface
type CatchesRepo interface {
	List() ([]*Catch, error)
	Get(id uuid.UUID) (*Catch, error)
	Create(c *Catch) (*Catch, error)
	Update(id uuid.UUID, c *Catch) (*Catch, error)
	Delete(id uuid.UUID) error

	ListByCompetitor(competitorID uuid.UUID) ([]*Catch, error)
	ListByTeam(teamID uuid.UUID) ([]*Catch, error)
}
