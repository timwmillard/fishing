package fishing

import (
	"time"

	"github.com/google/uuid"
)

// Event -
type Event struct {
	ID              uuid.UUID `firestore:"-"`
	Name            string    `firestore:"name,omitempty"`
	Slug            string    `firestore:"slug,omitempty"`
	CompetitionName string    `firestore:"competitionName,omitempty"`
	StartDate       time.Time `firestore:"startDate,omitempty"`
	EndDate         time.Time `firestore:"endDate,omitempty"`
}
