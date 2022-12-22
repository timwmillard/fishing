package http

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/timwmillard/fishing"
)

// CompetitorHandler -
type CompetitorHandler struct {
	repo   fishing.CompetitorRepo
	Router *mux.Router
	log    *log.Logger
	// errLog  *log.Logger
	// infoLog *log.Logger
}

// NewCompetitorsHandler -
func NewCompetitorHandler(repo fishing.CompetitorRepo) *CompetitorHandler {
	h := &CompetitorHandler{
		repo:   repo,
		Router: mux.NewRouter(),
		log:    log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	h.routes()
	return h
}

// ServeHTTP -
func (h *CompetitorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Router.ServeHTTP(w, r)
}

// List -
func (h *CompetitorHandler) List(w http.ResponseWriter, r *http.Request) {
	competitors, err := h.repo.List(r.Context())
	if err != nil {
		h.log.Printf("List Competitors: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(competitors)
}

// Get -
func (h *CompetitorHandler) Get(w http.ResponseWriter, r *http.Request) {
	competitor, err := h.repo.Get(r.Context(), fishing.HashID(mux.Vars(r)["id"]))
	if err != nil {
		h.log.Printf("Get Competitor: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(competitor)
}

// Create -
func (h *CompetitorHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request fishing.CreateCompetitorParams
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.Printf("Create Competitor Decode: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newCompetitor, err := h.repo.Create(r.Context(), request)
	if err != nil {
		h.log.Printf("Create Competitor: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newCompetitor)
}

// Update -
func (h *CompetitorHandler) Update(w http.ResponseWriter, r *http.Request) {
	var (
		request fishing.CreateCompetitorParams
		err     error
	)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.Printf("Update Competitor Decode: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedCompetitor, err := h.repo.Update(r.Context(), fishing.HashID(mux.Vars(r)["id"]), request)
	if err != nil {
		h.log.Printf("Update Competitor: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedCompetitor)
}

// Update -
func (h *CompetitorHandler) UpdatePartial(w http.ResponseWriter, r *http.Request) {
	var (
		request fishing.UpdateCompetitorParams
		err     error
	)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.Printf("Update Competitor Decode: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedCompetitor, err := h.repo.UpdatePartial(r.Context(), fishing.HashID(mux.Vars(r)["id"]), request)
	if err != nil {
		h.log.Printf("Update Competitor: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedCompetitor)
}

// Delete -
func (h *CompetitorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	err := h.repo.Delete(r.Context(), fishing.HashID(mux.Vars(r)["id"]))
	if err != nil {
		h.log.Printf("Delete Competitor: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
