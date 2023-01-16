package fishing

import "github.com/google/uuid"

// Team -
type Team struct {
	ID          uuid.UUID    `firestore:"-"`
	TeamNo      string       `firestore:"teamNo,omitempty"`
	Name        string       `firestore:"name,omitempty"`
	BoatRego    string       `firestore:"boatRego,omitempty"`
	Competitors []Competitor `firestore:"-"`
}

// TeamsRepo interface
type TeamsRepo interface {
	// List will list all teams without competitors
	List() ([]Team, error)

	// List will list all teams with competitors
	ListWithCompetitors() ([]Team, error)

	// Get a single team with all competitors
	Get(id uuid.UUID) (Team, error)

	// Create a empty team
	Create(c *Team) (Team, error)

	// Update team details
	Update(id uuid.UUID, c Team) (Team, error)

	// Delete a team
	Delete(id uuid.UUID) error

	// Competitors will list all competitors
	Competitors(teamID uuid.UUID) ([]Competitor, error)

	// Add single Competitor
	AddCompetitor(teamID uuid.UUID, competitorID int) error

	// Add multiple Competitors
	AddCompetitors(teamID uuid.UUID, competitorIDs []int) error

	// Remove single Competitor
	RemoveCompetitor(teamID uuid.UUID, competitorID int) error
}
