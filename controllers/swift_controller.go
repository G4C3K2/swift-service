package controllers

import (
	"log"
	"net/http"

	"github.com/G4C3K2/swift-service/services"
	"github.com/G4C3K2/swift-service/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection

func LoadSwiftData(c *gin.Context) {
	log.Println("Loading CSV...")
	records, err := utils.ParseCSV("Data.csv")
	if err != nil {
		log.Printf("CSV Parsing Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "CSV Parsing Error"})
		return
	}

	log.Printf("Loaded %d records from CSV", len(records))

	err = services.SaveSwiftEntries(records, Collection)
	if err != nil {
		log.Printf("Database write error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database write error"})
		return
	}

	log.Println("Data saved successfuly")
	c.JSON(http.StatusCreated, gin.H{"message": "Data saved successfuly"})
}

func GetSwiftCodeDetails(c *gin.Context) {
	swiftCode := c.Param("swiftCode")
	log.Printf("Searching for Swift Details: %s", swiftCode)

	result, err := services.GetSwiftCodeDetails(swiftCode, Collection)
	if err != nil {
		log.Printf("Swift Searching Error '%s': %v", swiftCode, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "SWIFT code not found"})
		return
	}

	log.Printf("Swift Details Found: %s", swiftCode)
	c.JSON(http.StatusOK, result)
}

func GetCountryISO2Details(c *gin.Context) {
	countryISO2 := c.Param("countryISO2code")
	log.Printf("Searching for Country ISO2 Code: %s", countryISO2)

	result, err := services.GetCountryISO2Details(countryISO2, Collection)
	if err != nil {
		log.Printf("Country ISO2 Code Found '%s': %v", countryISO2, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Country ISO2 Code not found"})
		return
	}

	log.Printf("Country ISO2 Code Found: %s", countryISO2)
	c.JSON(http.StatusOK, result)
}
