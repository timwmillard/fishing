package main

import (
	"fmt"
	"log"

	"github.com/timwmillard/fishing"
	"github.com/timwmillard/fishing/http"
	"github.com/timwmillard/fishing/memory"
)

func main() {

	log.Println("Starting Fishing Comp Server ...")

	// ctx := context.Background()

	compRepo := memory.NewCompetitorRepo()
	competitors := fishing.NewCompetitorService(compRepo)

	go func() {
		for e := range competitors.Events() {
			fmt.Printf("EVENT = %v\n", e.Message)
			e.Done()
		}
		println("Fishing go func()")
	}()

	// competitors.Create(ctx, fishing.Competitor{
	// 	CompetitorNo: "123",
	// 	Firstname:    "Tim",
	// 	Lastname:     "Millard",
	// 	Email:        "timwmillard@gmail.com",
	// 	Address1:     "75 Cassidy Lane",
	// 	Address2:     "",
	// 	Suburb:       "Koondrook",
	// 	State:        "VIC",
	// 	Postcode:     "3580",
	// 	Phone:        "",
	// 	Mobile:       "",
	// })

	// connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "?", "?", "localhost", "5432", "fishingcomp")
	// db, err := sql.Open("postgres", connectionString)
	// if err != nil {
	// 	return fmt.Errorf("database connection error: %v", err)
	// }
	// competitorRepo := postgres.NewCompetitorRepo(db)

	server := http.Server{
		CompetitorService: competitors,
	}

	err := server.ListenAndServe()
	log.Fatal(err)
}
