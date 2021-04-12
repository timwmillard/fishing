package main

import (
	"fishing/http"
	"log"
)

var server http.Server

func main() {

	log.Println("Starting Fishing Comp Server ...")

	err := server.ListenAndServe()
	log.Fatal(err)
}
