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
		handleError(w, err, "list competitor")
		return
	}
	json.NewEncoder(w).Encode(competitors)
}

// Get -
func (h *CompetitorHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := fishing.NewHashID(mux.Vars(r)["id"])
	if err != nil {
		handleError(w, err, "get competitor")
		return
	}
	competitor, err := h.repo.Get(r.Context(), id)
	if err != nil {
		handleError(w, err, "get competitor")
		return
	}
	json.NewEncoder(w).Encode(competitor)
}

// Create -
func (h *CompetitorHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request fishing.CreateCompetitorParams
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleJSONDecodeError(w, err, "create competitor")
		return
	}
	newCompetitor, err := h.repo.Create(r.Context(), request)
	if err != nil {
		handleError(w, err, "create competitor")
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

	id, err := fishing.NewHashID(mux.Vars(r)["id"])
	if err != nil {
		handleError(w, err, "update competitor")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleJSONDecodeError(w, err, "update competitor")
		return
	}

	updatedCompetitor, err := h.repo.Update(r.Context(), id, request)
	if err != nil {
		handleError(w, err, "update competitor")
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

	id, err := fishing.NewHashID(mux.Vars(r)["id"])
	if err != nil {
		handleError(w, err, "partial update competitor")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleJSONDecodeError(w, err, "partial update competitor")
		return
	}

	updatedCompetitor, err := h.repo.UpdatePartial(r.Context(), id, request)
	if err != nil {
		handleError(w, err, "partial update competitor")
		return
	}
	json.NewEncoder(w).Encode(updatedCompetitor)
}

// Delete -
func (h *CompetitorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := fishing.NewHashID(mux.Vars(r)["id"])
	if err != nil {
		handleError(w, err, "update competitor")
		return
	}
	err = h.repo.Delete(r.Context(), id)
	if err != nil {
		handleError(w, err, "delete competitor")
		return
	}
	w.WriteHeader(http.StatusOK)
}
