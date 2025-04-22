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
	log.Printf("SaveSwiftEntries: Start - number of input records: %d\n", len(data))
	var entries []interface{}

	for i, record := range data {
		swift := record["SWIFT CODE"]
		code := strings.ToUpper(record["COUNTRY ISO2 CODE"])
		name := strings.ToUpper(record["NAME"])
		address := record["ADDRESS"]

		log.Printf("SaveSwiftEntries: Processing record #%d: SWIFT=%s, CODE=%s, NAME=%s\n", i+1, swift, code, name)

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

	log.Printf("SaveSwiftEntries: Ready to save %d records\n", len(entries))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertMany(ctx, entries)
	if err != nil {
		log.Printf("SaveSwiftEntries: Error saving to Mongo: %v\n", err)
		return err
	}
	log.Println("SaveSwiftEntries: Save completed successfully.")
	return nil
}

func GetSwiftCodeDetails(swiftCode string, collection *mongo.Collection) (*models.SwiftCodeResponse, error) {
	log.Printf("GetSwiftCodeDetails: Fetching details for SWIFT code = %s\n", swiftCode)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mainEntry, err := repository.FindBySwiftCode(ctx, collection, swiftCode)
	if err != nil {
		log.Printf("GetSwiftCodeDetails: Main entry not found: %v\n", err)
		return nil, err
	}

	log.Printf("GetSwiftCodeDetails: Found main entry: %+v\n", mainEntry)

	response := &models.SwiftCodeResponse{
		Address:       mainEntry.Address,
		BankName:      mainEntry.Name,
		CountryISO2:   mainEntry.CountryCode,
		CountryName:   mainEntry.CountryName,
		IsHeadquarter: mainEntry.IsHeadquarter,
		SwiftCode:     mainEntry.SwiftCode,
	}

	if mainEntry.IsHeadquarter {
		log.Printf("GetSwiftCodeDetails: This is a headquarter, searching for branches for HQ = %s\n", mainEntry.SwiftCode)

		branches, err := repository.FindBranchesByHqCode(ctx, collection, mainEntry.SwiftCode)
		if err != nil {
			log.Printf("GetSwiftCodeDetails: Error fetching branches: %v\n", err)
			return nil, err
		}

		log.Printf("GetSwiftCodeDetails: Found %d branches\n", len(branches))

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

	log.Printf("GetSwiftCodeDetails: Returning response: %+v\n", response)
	return response, nil
}

func GetCountryISO2Details(countryISO2 string, collection *mongo.Collection) (*models.CountryISO2CodeResponse, error) {
	log.Printf("GetCountryISO2Details: Fetching details for countryISO2 = %s\n", countryISO2)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mainEntry, err := repository.FindByCountryCode(ctx, collection, countryISO2)
	if err != nil {
		log.Printf("GetCountryISO2Details: Country not found: %v\n", err)
		return nil, err
	}

	log.Printf("GetCountryISO2Details: Found country: %+v\n", mainEntry)

	response := &models.CountryISO2CodeResponse{
		CountryISO2: mainEntry.CountryCode,
		CountryName: mainEntry.CountryName,
	}

	banks, err := repository.FindBanksByCountry(ctx, collection, mainEntry.CountryCode)
	if err != nil {
		log.Printf("GetCountryISO2Details: Error fetching banks: %v\n", err)
		return nil, err
	}

	log.Printf("GetCountryISO2Details: Found %d banks\n", len(banks))

	for _, b := range banks {
		response.SwiftCodes = append(response.SwiftCodes, models.SwiftBranch{
			Address:       b.Address,
			BankName:      b.Name,
			CountryISO2:   b.CountryCode,
			IsHeadquarter: b.IsHeadquarter,
			SwiftCode:     b.SwiftCode,
		})
	}

	log.Printf("GetCountryISO2Details: Returning response: %+v\n", response)
	return response, nil
}

func CreateSwiftEntry(entry *models.SwiftEntry, collection *mongo.Collection) error {
	err := repository.InsertSwiftEntry(entry, collection)
	return err
}
