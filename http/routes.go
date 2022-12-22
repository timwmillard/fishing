package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) routes() {
	s.router = mux.NewRouter()

	// Middleware
	s.router.Use(jsonMiddleware)
	s.router.Use(corsMiddleware)

	// Routes
	s.router.HandleFunc("/", index)
	s.router.HandleFunc("/competitors", s.competitorsHandler.List).Methods(http.MethodGet)           // Get all contacts
	s.router.HandleFunc("/competitors/{id}", s.competitorsHandler.Get).Methods(http.MethodGet)       // Get contact
	s.router.HandleFunc("/competitors", s.competitorsHandler.Create).Methods(http.MethodPost)        // Create a contact
	s.router.HandleFunc("/competitors/{id}", s.competitorsHandler.Update).Methods(http.MethodPut)    // Update a contact
	s.router.HandleFunc("/competitors/{id}", s.competitorsHandler.Delete).Methods(http.MethodDelete) // Update a contact

}
