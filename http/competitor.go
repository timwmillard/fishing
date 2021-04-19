package http

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/timwmillard/fishing"
)

// CompetitorHandler -
type CompetitorHandler struct {
	repo   fishing.CompetitorRepo
	router *mux.Router
	log    *log.Logger
	// errLog  *log.Logger
	// infoLog *log.Logger
}

// NewCompetitorsHandler -
func NewCompetitorHandler(repo fishing.CompetitorRepo) *CompetitorHandler {
	return &CompetitorHandler{
		repo:   repo,
		router: mux.NewRouter(),
		log:    log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// ServeHTTP -
func (c *CompetitorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.router.ServeHTTP(w, r)
}

// List -
func (c *CompetitorHandler) List(w http.ResponseWriter, r *http.Request) {
	competitors, err := c.repo.List(r.Context())
	if err != nil {
		c.log.Printf("List Competitors: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(competitors)
}

// Get -
func (c *CompetitorHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		c.log.Printf("Get Competitor: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	competitor, err := c.repo.Get(r.Context(), id)
	if err != nil {
		c.log.Printf("Get Competitor: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(competitor)
}

// Create -
func (c *CompetitorHandler) Create(w http.ResponseWriter, r *http.Request) {
	var requestCompetitor fishing.Competitor
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestCompetitor)
	if err != nil {
		c.log.Printf("Create Competitor Decode: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newCompetitor, err := c.repo.Create(r.Context(), requestCompetitor)
	if err != nil {
		c.log.Printf("Create Competitor: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(newCompetitor)
}

// Update -
func (c *CompetitorHandler) Update(w http.ResponseWriter, r *http.Request) {
	var (
		requestCompetitor fishing.Competitor
		err               error
	)

	requestCompetitor.ID, err = uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		c.log.Printf("Update Competitor: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&requestCompetitor)
	if err != nil {
		c.log.Printf("Update Competitor Decode: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedCompetitor, err := c.repo.Update(r.Context(), requestCompetitor)
	if err != nil {
		c.log.Printf("Update Competitor: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(updatedCompetitor)
}

// Delete -
func (c *CompetitorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		c.log.Printf("Update Competitor: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.repo.Delete(r.Context(), id)
	if err != nil {
		c.log.Printf("Delete Competitor: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)

}
