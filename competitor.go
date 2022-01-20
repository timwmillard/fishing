package fishing

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// Common Errors
var (
	ErrCompetitorNotFound = errors.New("competitor not found")
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
	// EventID      uuid.UUID `json:"event_id,omitempty" db:"event_id"`
}

// CompetitorsRepo interface for competitors repository.
type CompetitorRepo interface {
	List(ctx context.Context) ([]Competitor, error)
	Get(ctx context.Context, id uuid.UUID) (Competitor, error)
	Create(ctx context.Context, c Competitor) (Competitor, error)
	Update(ctx context.Context, c Competitor) (Competitor, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
type CompetitorService struct {
	repo   CompetitorRepo
	events chan CompetitorEvent
}

func NewCompetitorService(repo CompetitorRepo) *CompetitorService {
	return &CompetitorService{
		repo:   repo,
		events: make(chan CompetitorEvent),
	}
}

func (cs *CompetitorService) Events() <-chan CompetitorEvent {
	return cs.events
}

func (cs *CompetitorService) List(ctx context.Context) ([]Competitor, error) {
	return cs.repo.List(ctx)
}

func (cs *CompetitorService) Get(ctx context.Context, id uuid.UUID) (Competitor, error) {
	return cs.repo.Get(ctx, id)
}

func (cs *CompetitorService) Create(ctx context.Context, c Competitor) (Competitor, error) {
	comp, err := cs.repo.Create(ctx, c)
	if err != nil {
		return Competitor{}, err
	}
	cs.sendEvent(CompetitorEvent{
		EventType:       CreateCompetitorEvent,
		AfterCompetitor: comp,
		Message:         fmt.Sprintf("Created new competitor %v", comp.ID),
	})
	return comp, err
}

func (cs *CompetitorService) Update(ctx context.Context, c Competitor) (Competitor, error) {
	comp, err := cs.repo.Update(ctx, c)
	if err != nil {
		return Competitor{}, err
	}
	cs.sendEvent(CompetitorEvent{
		EventType:        UpdateCompetitorEvent,
		BeforeCompetitor: c,
		AfterCompetitor:  comp,
		Message:          fmt.Sprintf("Updated competitor %v", c.ID),
	})
	return comp, err
}

func (cs *CompetitorService) Delete(ctx context.Context, id uuid.UUID) error {
	err := cs.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	cs.sendEvent(CompetitorEvent{
		EventType: DeleteCompetitorEvent,
		Message:   fmt.Sprintf("Deleted competitor %v", id),
	})
	return nil
}

// CompetitorEventType
type CompetitorEventType int

const (
	CreateCompetitorEvent CompetitorEventType = iota
	UpdateCompetitorEvent
	DeleteCompetitorEvent
)

type CompetitorEvent struct {
	EventType        CompetitorEventType
	BeforeCompetitor Competitor
	AfterCompetitor  Competitor
	Message          string
	err              chan error
}

func (e CompetitorEvent) Done() {
	e.err <- nil
}

func (e CompetitorEvent) Error(err error) {
	e.err <- err
}

// sendEvent sends a non blocking event onto the events channel.
func (cs *CompetitorService) sendEvent(e CompetitorEvent) error {
	e.err = make(chan error)
	select {
	case cs.events <- e:
	default:
	}
	return <-e.err
}
