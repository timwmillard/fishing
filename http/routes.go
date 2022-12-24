package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *CompetitorHandler) routes() {
	h.Router = mux.NewRouter()

	// Middleware
	h.Router.Use(auth("test"))
	h.Router.Use(jsonMiddleware)
	h.Router.Use(corsMiddleware)

	// Routes
	h.Router.HandleFunc("/", index)
	h.Router.HandleFunc("/health", health)
	h.Router.HandleFunc("/competitors", h.List).Methods(http.MethodGet)                 // Get all competitors
	h.Router.HandleFunc("/competitors", h.Create).Methods(http.MethodPost)              // Create a competitors
	h.Router.HandleFunc("/competitors/{id}", h.Get).Methods(http.MethodGet)             // Get a competitors
	h.Router.HandleFunc("/competitors/{id}", h.Update).Methods(http.MethodPut)          // Update a competitors
	h.Router.HandleFunc("/competitors/{id}", h.UpdatePartial).Methods(http.MethodPatch) // Partial update a competitors
	h.Router.HandleFunc("/competitors/{id}", h.Delete).Methods(http.MethodDelete)       // Delete a competitors
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Fishing Comp App")
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("Health Check")
	w.WriteHeader(http.StatusOK)
}
