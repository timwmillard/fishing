package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/timwmillard/fishing/http"
	"github.com/timwmillard/fishing/postgres"
)

func main() {
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)

	godotenv.Load()

	portEnv := os.Getenv("PORT")
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		log.Fatalf("Port must be an integer: %v", err)
	}
	dbHost := os.Getenv("PGHOST")
	dbPort := os.Getenv("PGPORT")
	dbName := os.Getenv("PGDATABASE")
	dbUser := os.Getenv("PGUSER")
	dbPassword := os.Getenv("PGPASSWORD")

	// Setup database
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	log.Printf("Connecting to Postgres at %s", dbURL)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Unable to open PostgreSQL database: %v", err)
	}
	log.Println("Postgres database connection successful")

	competitorRepo := &postgres.CompetitorRepo{DB: db}
	competitorHandler := http.NewCompetitorHandler(competitorRepo)

	server := http.NewServer(port, competitorHandler)
	server.Start()

	defer func(server *http.Server) {
		err := server.Stop()
		if err != nil {
			log.Fatal(err)
		}
	}(server)

	// Wait for CTRL-C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	// We block here until a CTRL-C / SigInt is received
	// Once received, we exit and the server is cleaned up
	<-sigChan
	cancel()

}
