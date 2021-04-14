package main

import (
	"log"

	"github.com/timwmillard/fishing/http"
)

var server http.Server

func main() {

	log.Println("Starting Fishing Comp Server ...")

	err := server.ListenAndServe()
	log.Fatal(err)
}
