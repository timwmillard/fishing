package fishing

import "github.com/google/uuid"

// Species - a fish species
// Examples
// Murray Cod
// Yellow Belly
type Species struct {
	UUID           uuid.UUID
	Slug           string
	CommonName     string
	ScientificName string
	PhotoURL       string
}

// SpeciesRepo interface
type SpeciesRepo interface {
	List() ([]*Species, error)
	Get(id uuid.UUID) (*Species, error)
	Create(c *Species) (*Species, error)
	Update(id uuid.UUID, c *Species) (*Species, error)
	Delete(id uuid.UUID) error
}
