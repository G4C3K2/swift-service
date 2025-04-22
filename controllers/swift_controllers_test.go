package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/G4C3K2/swift-service/models"
	"github.com/gin-gonic/gin"
)

// prosta pomocnicza funkcja do inicjalizacji routera
func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swift/:swiftCode", GetSwiftCodeDetails)
	router.GET("/country/:countryISO2code", GetCountryISO2Details)
	router.POST("/swift", AddSwiftCode)
	router.DELETE("/swift/:swiftCode", DeleteSwiftCode)

	return router
}

func TestAddSwiftCode_InvalidInput(t *testing.T) {
	router := setupRouter()

	body := `{"SwiftCode": "INVALID", "CountryISO2": "PL"}`
	req, _ := http.NewRequest("POST", "/swift", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", resp.Code)
	}
}

func TestAddSwiftCode_ValidInput(t *testing.T) {
	router := setupRouter()

	// zakładamy że nie dojdzie do faktycznego zapisu w bazie
	payload := models.AddSwiftCodeRequest{
		SwiftCode:     "TESTPLPAXXX",
		CountryISO2:   "PL",
		CountryName:   "Poland",
		BankName:      "Test Bank",
		Address:       "Test Street 123",
		IsHeadquarter: true,
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/swift", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.Code)
	}
}

func TestDeleteSwiftCode_MissingParam(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/swift/", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 for missing param, got %d", resp.Code)
	}
}

func TestGetSwiftCodeDetails_NotFound(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/swift/NOSUCHCODE", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", resp.Code)
	}
}
