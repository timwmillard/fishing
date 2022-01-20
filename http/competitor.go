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
	service *fishing.CompetitorService
	router  *mux.Router
	log     *log.Logger
	// errLog  *log.Logger
	// infoLog *log.Logger
}

// NewCompetitorsHandler -
func NewCompetitorHandler(service *fishing.CompetitorService) *CompetitorHandler {
	return &CompetitorHandler{
		service: service,
		router:  mux.NewRouter(),
		log:     log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// ServeHTTP -
func (c *CompetitorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.router.ServeHTTP(w, r)
}

// List -
func (c *CompetitorHandler) List(w http.ResponseWriter, r *http.Request) {
	competitors, err := c.service.List(r.Context())
	if err != nil {
		c.log.Printf("List Competitors: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(competitors)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Get -
func (c *CompetitorHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		c.log.Printf("Get Competitor: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	competitor, err := c.service.Get(r.Context(), id)
	if err != nil {
		c.log.Printf("Get Competitor: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(competitor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
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
	newCompetitor, err := c.service.Create(r.Context(), requestCompetitor)
	if err != nil {
		c.log.Printf("Create Competitor: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(newCompetitor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
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

	updatedCompetitor, err := c.service.Update(r.Context(), requestCompetitor)
	if err != nil {
		c.log.Printf("Update Competitor: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(updatedCompetitor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Delete -
func (c *CompetitorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		c.log.Printf("Update Competitor: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.service.Delete(r.Context(), id)
	if err != nil {
		c.log.Printf("Delete Competitor: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
