package http

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/timwmillard/fishing"
)

// Server -
type Server struct {
	DB     *sql.DB
	router *mux.Router

	CompetitorService *fishing.CompetitorService

	// Handlers
	competitorsHandler *CompetitorHandler
}

// ListenAndServe -
func (s *Server) ListenAndServe() error {
	var err error

	s.competitorsHandler = NewCompetitorHandler(s.CompetitorService)

	// Setup Routing
	s.routes()

	fmt.Printf("Listing on port 6000\n")
	err = http.ListenAndServe(":6000", s.router)
	if err != nil {
		return err
	}
	return nil
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Fishing Comp App")
}
