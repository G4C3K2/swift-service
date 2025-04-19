package utils

import (
	"log"
	"strings"

	"github.com/G4C3K2/swift-service/models"
)

func ParseSheetData(values [][]interface{}) []models.SwiftEntry {
	var entries []models.SwiftEntry

	for i, row := range values {
		if i == 0 {
			continue // pomiń nagłówki
		}

		if len(row) < 6 {
			log.Printf("Pominięto wiersz %d - za mało kolumn", i+1)
			continue
		}

		entry := models.SwiftEntry{
			SwiftCode:     toString(row[0]),
			BankName:      toString(row[1]),
			Address:       toString(row[2]),
			CountryISO2:   toString(row[3]),
			CountryName:   toString(row[4]),
			IsHeadquarter: toBool(row[5]),
		}

		if len(row) >= 7 {
			entry.HeadquarterCode = toString(row[6])
		}

		entries = append(entries, entry)
	}

	return entries
}

func toString(val interface{}) string {
	str, ok := val.(string)
	if !ok {
		return ""
	}
	return strings.TrimSpace(str)
}

func toBool(val interface{}) bool {
	str := toString(val)
	return strings.ToLower(str) == "true" || str == "1" || str == "yes"
}
