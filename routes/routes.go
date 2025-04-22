package routes

import (
	"log"

	"github.com/G4C3K2/swift-service/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(router *gin.Engine, collection *mongo.Collection) {
	controllers.Collection = collection

	log.Println("Routing...")

	router.GET("/ping", func(c *gin.Context) {
		log.Println("PING received!")
		c.JSON(200, gin.H{"message": "pong"})
	})

	swiftGroup := router.Group("/swift-codes")
	{
		swiftGroup.POST("/load", controllers.LoadSwiftData)

		swiftGroup.GET("/:swiftCode", controllers.GetSwiftCodeDetails)

		swiftGroup.GET("/country/:countryISO2code", controllers.GetCountryISO2Details)

		swiftGroup.POST("/", controllers.AddSwiftCode)
	}

	log.Println("Routing finished.")
}
