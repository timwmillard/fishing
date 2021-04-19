package http

import "github.com/gorilla/mux"

func (s *Server) routes() {
	s.router = mux.NewRouter()

	// Middleware
	s.router.Use(jsonMiddleware)
	s.router.Use(corsMiddleware)

	// Routes
	s.router.HandleFunc("/", index)
	s.router.HandleFunc("/competitors", s.competitorsHandler.List).Methods("GET")           // Get all contacts
	s.router.HandleFunc("/competitors/{id}", s.competitorsHandler.Get).Methods("GET")       // Get contact
	s.router.HandleFunc("/competitors", s.competitorsHandler.Create).Methods("POST")        // Create a contact
	s.router.HandleFunc("/competitors/{id}", s.competitorsHandler.Update).Methods("PUT")    // Update a contact
	s.router.HandleFunc("/competitors/{id}", s.competitorsHandler.Delete).Methods("DELETE") // Update a contact

}
