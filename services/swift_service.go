package services

import (
	"context"
	"log"
	"os"

	"github.com/G4C3K2/swift-service/repository"
	"github.com/G4C3K2/swift-service/utils"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func FetchAndStoreFromSheet(sheetID, readRange string) {
	ctx := context.Background()
	creds := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(creds))
	if err != nil {
		log.Fatalf("Failed to create Sheets service: %v", err)
	}

	resp, err := srv.Spreadsheets.Values.Get(sheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	entries := utils.ParseSheetData(resp.Values)
	repository.InsertMany(entries)
}
