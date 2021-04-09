package http

import (
	"context"
	"fishing"
	"fishing/memory"
	"fishing/postgres"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"database/sql"

	"gorm.io/gorm"
)

// Server -
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// ListenAndServe -
func (s *Server) ListenAndServe() error {
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "?", "?", "localhost", "5432", "fishingcomp")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("database connection error: %v", err)
	}

	// var competitorsRepo fishing.CompetitorsRepo

	// competitorsRepo := memory.NewCompetitorsRepo()
	competitorsRepo := postgres.NewCompetitorsRepo(db)

	// createDummyData(competitorsRepo)

	competitorsHandler := NewCompetitorsHandler(competitorsRepo)

	// Setup Routing
	s.Router = mux.NewRouter()

	s.Router.HandleFunc("/", index)
	s.Router.HandleFunc("/competitors", competitorsHandler.List).Methods("GET")           // Get all contacts
	s.Router.HandleFunc("/competitors/{id}", competitorsHandler.Get).Methods("GET")       // Get contact
	s.Router.HandleFunc("/competitors", competitorsHandler.Create).Methods("POST")        // Create a contact
	s.Router.HandleFunc("/competitors/{id}", competitorsHandler.Update).Methods("PUT")    // Update a contact
	s.Router.HandleFunc("/competitors/{id}", competitorsHandler.Delete).Methods("DELETE") // Update a contact

	fmt.Printf("Listing on port 8080\n")
	err = http.ListenAndServe(":8080", s.Router)
	if err != nil {
		return err
	}
	return nil
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Fishing Comp App")
}

//lint:ignore U1000 unused
func createDummyData(repo *memory.CompetitorsRepo) {
	ctx := context.Background()
	competitor := fishing.Competitor{
		Firstname: "Tim",
		Lastname:  "Millard",
	}
	repo.Create(ctx, competitor)
	competitor = fishing.Competitor{
		Firstname: "John",
		Lastname:  "Smith",
	}
	repo.Create(ctx, competitor)
}
