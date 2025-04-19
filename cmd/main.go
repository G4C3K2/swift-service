package main

import (
	"log"
	"net/http"

	"github.com/G4C3K2/swift-service/routes"
)

func main() {
	// Zainicjalizuj router (to musisz mieć w routes/setup.go lub czymś podobnym)
	router := routes.SetupRouter()

	log.Println("Starting HTTP server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
