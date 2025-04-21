package services

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/G4C3K2/swift-service/models"
	"github.com/G4C3K2/swift-service/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveSwiftEntries(data []map[string]string, collection *mongo.Collection) error {
	var entries []interface{}

	for _, record := range data {
		swift := record["SWIFT CODE"]
		code := strings.ToUpper(record["COUNTRY ISO2 CODE"])
		name := strings.ToUpper(record["NAME"])
		address := record["ADDRESS"]

		var addrPtr *string
		if strings.TrimSpace(address) != "" {
			addrPtr = &address
		}

		var hqCode *string
		if !strings.HasSuffix(swift, "XXX") {
			hq := swift[:len(swift)-3] + "XXX"
			hqCode = &hq
		}

		entry := models.SwiftEntry{
			SwiftCode:     swift,
			CodeType:      record["CODE TYPE"],
			Name:          name,
			Address:       addrPtr,
			TownName:      record["TOWN NAME"],
			CountryCode:   code,
			CountryName:   record["COUNTRY NAME"],
			TimeZone:      record["TIME ZONE"],
			IsHeadquarter: strings.HasSuffix(swift, "XXX"),
			HqCode:        hqCode,
		}

		entries = append(entries, entry)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertMany(ctx, entries)
	if err != nil {
		log.Printf("Record entry failed: %v\n", err)
		return err
	}
	return nil
}

func GetSwiftCodeDetails(swiftCode string, collection *mongo.Collection) (*models.SwiftCodeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mainEntry, err := repository.FindBySwiftCode(ctx, collection, swiftCode)
	if err != nil {
		return nil, err
	}

	response := &models.SwiftCodeResponse{
		Address:       mainEntry.Address,
		BankName:      mainEntry.Name,
		CountryISO2:   mainEntry.CountryCode,
		CountryName:   mainEntry.CountryName,
		IsHeadquarter: mainEntry.IsHeadquarter,
		SwiftCode:     mainEntry.SwiftCode,
	}

	if mainEntry.IsHeadquarter {
		branches, err := repository.FindBranchesByHqCode(ctx, collection, mainEntry.SwiftCode)
		if err != nil {
			return nil, err
		}

		for _, b := range branches {
			response.Branches = append(response.Branches, models.SwiftBranch{
				Address:       b.Address,
				BankName:      b.Name,
				CountryISO2:   b.CountryCode,
				IsHeadquarter: b.IsHeadquarter,
				SwiftCode:     b.SwiftCode,
			})
		}
	}

	return response, nil
}
