package main

import (
	"fishing/http"
	"log"
)

var server http.Server

func main() {

	err := server.ListenAndServe()
	log.Fatal(err)
}
