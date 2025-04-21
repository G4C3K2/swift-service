package controllers

import (
	"net/http"

	"github.com/G4C3K2/swift-service/services"
	"github.com/G4C3K2/swift-service/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection

func LoadSwiftData(c *gin.Context) {
	records, err := utils.ParseCSV("Data.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd parsowania CSV"})
		return
	}

	err = services.SaveSwiftEntries(records, Collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd zapisu do bazy"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Dane załadowane poprawnie"})
}

func GetSwiftCodeDetails(c *gin.Context) {
	swiftCode := c.Param("swiftCode")

	result, err := services.GetSwiftCodeDetails(swiftCode, Collection)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SWIFT code not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}
