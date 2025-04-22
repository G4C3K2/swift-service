package main

import (
	"log"

	"github.com/G4C3K2/swift-service/config"
	"github.com/G4C3K2/swift-service/controllers" // Zaimportuj pakiet controllers
	"github.com/G4C3K2/swift-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	collection := config.DatabaseConnection()

	// Przypisz instancję *mongo.Collection do globalnej zmiennej w controllers
	controllers.Collection = collection
	log.Println("MongoDB Collection initialized in main.go")

	router := gin.Default()

	routes.SetupRoutes(router, collection) // Przekaż tę samą instancję do routerów

	log.Println("Routing completed, starting server")

	log.Println("Listening on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Starting server error:", err)
	}
}
