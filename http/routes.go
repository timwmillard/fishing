package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *CompetitorHandler) routes() {
	router := mux.NewRouter()
	h.Router = router

	// Middleware
	router.Use(jsonMiddleware)
	router.Use(corsMiddleware)
	router.HandleFunc("/", index)
	router.HandleFunc("/health", health)

	compRouter := router.PathPrefix("/competitors").Subrouter()

	compRouter.Use(auth("test"))

	// Routes
	compRouter.HandleFunc("", h.List).Methods(http.MethodGet)                 // Get all competitors
	compRouter.HandleFunc("", h.Create).Methods(http.MethodPost)              // Create a competitors
	compRouter.HandleFunc("/{id}", h.Get).Methods(http.MethodGet)             // Get a competitors
	compRouter.HandleFunc("/{id}", h.Update).Methods(http.MethodPut)          // Update a competitors
	compRouter.HandleFunc("/{id}", h.UpdatePartial).Methods(http.MethodPatch) // Partial update a competitors
	compRouter.HandleFunc("/{id}", h.Delete).Methods(http.MethodDelete)       // Delete a competitors
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Fishing Comp App")
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("Health Check")
	w.WriteHeader(http.StatusOK)
}
