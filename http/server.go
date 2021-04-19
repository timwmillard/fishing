package http

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/timwmillard/fishing/postgres"
)

// Server -
type Server struct {
	DB     *sql.DB
	router *mux.Router

	// Handlers
	competitorsHandler *CompetitorsHandler
}

// ListenAndServe -
func (s *Server) ListenAndServe() error {
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "?", "?", "localhost", "5432", "fishingcomp")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("database connection error: %v", err)
	}

	competitorsRepo := postgres.NewCompetitorsRepo(db)
	s.competitorsHandler = NewCompetitorsHandler(competitorsRepo)

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
