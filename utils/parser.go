package utils

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/G4C3K2/swift-service/models"
)

func LoadData(fileName string, collection *mongo.Collection) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Cannot read the file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading records:", err)
	}

	if len(records) < 2 {
		log.Println("Brak danych w pliku lub tylko nagłówek")
		return
	}

	header := records[0]
	indexMap := map[string]int{}
	for i, col := range header {
		indexMap[strings.ToUpper(strings.TrimSpace(col))] = i
	}

	get := func(record []string, field string) string {
		idx, ok := indexMap[strings.ToUpper(field)]
		if !ok || idx >= len(record) {
			return ""
		}
		return strings.TrimSpace(record[idx])
	}

	const batchSize = 100
	var batch []interface{}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	for i, record := range records[1:] {
		swift := get(record, "SWIFT CODE")
		isHQ := strings.HasSuffix(swift, "XXX")

		addr := get(record, "ADDRESS")
		var addressPtr *string
		if addr != "" {
			addressPtr = &addr
		}

		entry := models.SwiftEntry{
			SwiftCode:     swift,
			CodeType:      get(record, "CODE TYPE"),
			Name:          get(record, "NAME"),
			Address:       addressPtr,
			TownName:      get(record, "TOWN NAME"),
			CountryCode:   get(record, "COUNTRY ISO2 CODE"),
			CountryName:   get(record, "COUNTRY NAME"),
			TimeZone:      get(record, "TIME ZONE"),
			IsHeadquarter: isHQ,
		}

		batch = append(batch, entry)

		// Wysyłaj paczką co 100 lub na końcu
		if len(batch) == batchSize || i == len(records[1:])-1 {
			_, err := collection.InsertMany(ctx, batch)
			if err != nil {
				log.Println("Błąd przy batchowym dodawaniu:", err)
			} else {
				fmt.Printf("Dodano batch %d rekordów\n", len(batch))
			}
			batch = nil
		}
	}
}
