package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
	"github.com/timwmillard/fishing/http"
	"github.com/timwmillard/fishing/postgres"
)

var port = 9600

var (
	dbName     = "fishingcomp"
	dbUser     = "root"
	dbPassword = "fish"
	dbPort     = 5555
	dbHost     = "localhost"
)

func main() {
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)

	// Setup database
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	println(dbURI)
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	competitorRepo := &postgres.CompetitorRepo{DB: db}
	competitorHandler := http.NewCompetitorHandler(competitorRepo)

	server := http.NewServer(port, competitorHandler.Router)
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
