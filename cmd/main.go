package main

import (
	"log"

	"github.com/G4C3K2/swift-service/config"
	"github.com/G4C3K2/swift-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	collection := config.DatabaseConnection()

	router := gin.Default()

	routes.SetupRoutes(router, collection)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Swift service is running...",
		})
	})

	log.Println("Listening on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Błąd uruchomienia serwera:", err)
	}
}
