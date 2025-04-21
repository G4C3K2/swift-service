package routes

import (
	"github.com/G4C3K2/swift-service/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(router *gin.Engine, collection *mongo.Collection) {
	controllers.Collection = collection

	swiftGroup := router.Group("/swift-codes")
	{
		swiftGroup.POST("/load", controllers.LoadSwiftData)
		swiftGroup.GET("/:swiftCode", controllers.GetSwiftCodeDetails)
	}
}
