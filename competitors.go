package fishing

import "github.com/google/uuid"

// Competitor -
type Competitor struct {
	ID           uuid.UUID `json:"id,omitempty" db:"id"`
	CompetitorNo *int      `json:"competitor_no,omitempty" db:"competitor_no"`
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
	EventID      *int      `json:"event_id,omitempty" db:"event_id"`
}

// CompetitorsRepo interface
type CompetitorsRepo interface {
	List() ([]*Competitor, error)
	Get(id uuid.UUID) (*Competitor, error)
	// Find(query string)
	Create(c *Competitor) (*Competitor, error)
	Update(id uuid.UUID, c *Competitor) (*Competitor, error)
	Delete(id uuid.UUID) error
}
