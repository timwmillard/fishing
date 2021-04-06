package http

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"fishing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// CompetitorsHandler -
type CompetitorsHandler struct {
	repo   fishing.CompetitorsRepo
	router *mux.Router
	log    *log.Logger
	// errLog  *log.Logger
	// infoLog *log.Logger
}

// NewCompetitorsHandler -
func NewCompetitorsHandler(repo fishing.CompetitorsRepo) *CompetitorsHandler {

	h := &CompetitorsHandler{
		repo:   repo,
		router: mux.NewRouter(),
		log:    log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
		// errLog:  log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
		// infoLog: log.New(os.Stderr, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	h.router.HandleFunc("/competitors", h.List).Methods("GET")           // Get all contacts
	h.router.HandleFunc("/competitors/{id}", h.Get).Methods("GET")       // Get contact
	h.router.HandleFunc("/competitors", h.Create).Methods("POST")        // Create a contact
	h.router.HandleFunc("/competitors/{id}", h.Update).Methods("PUT")    // Update a contact
	h.router.HandleFunc("/competitors/{id}", h.Delete).Methods("DELETE") // Update a contact

	return h
}

// ServeHTTP -
func (c *CompetitorsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.router.ServeHTTP(w, r)
}

// List -
func (c *CompetitorsHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	competitors, err := c.repo.List()
	if err != nil {
		c.log.Printf("List Competitors: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(competitors)
}

// Get -
func (c *CompetitorsHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		c.log.Printf("Get Competitor: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	competitor, err := c.repo.Get(id)
	if err != nil {
		c.log.Printf("Get Competitor: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(competitor)
}

// Create -
func (c *CompetitorsHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var requestCompetitor fishing.Competitor
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestCompetitor)
	if err != nil {
		c.log.Printf("Create Competitor Decode: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newCompetitor, err := c.repo.Create(&requestCompetitor)
	if err != nil {
		c.log.Printf("Create Competitor: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(newCompetitor)
}

// Update -
func (c *CompetitorsHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		c.log.Printf("Update Competitor: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var requestCompetitor fishing.Competitor
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&requestCompetitor)
	if err != nil {
		c.log.Printf("Update Competitor Decode: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedCompetitor, err := c.repo.Update(id, &requestCompetitor)
	if err != nil {
		c.log.Printf("Update Competitor: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(updatedCompetitor)
}

// Delete -
func (c *CompetitorsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		c.log.Printf("Update Competitor: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.repo.Delete(id)
	if err != nil {
		c.log.Printf("Delete Competitor: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)

}
