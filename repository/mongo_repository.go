package repository

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/G4C3K2/swift-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertMany(entries []models.SwiftEntry) {
	mongoURI := os.Getenv("MONGO_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("swift").Collection("entries")

	var docs []interface{}
	for _, entry := range entries {
		docs = append(docs, entry)
	}

	_, err = collection.InsertMany(ctx, docs)
	if err != nil {
		log.Fatalf("Error saving to Mongo: %v", err)
	} else {
		log.Printf("Successfully saved %d documents.\n", len(docs))
	}
}

func FindBySwiftCode(ctx context.Context, collection *mongo.Collection, swiftCode string) (*models.SwiftEntry, error) {
	log.Printf("FindBySwiftCode: Searching for document with SWIFT code = %s\n", swiftCode)

	var result models.SwiftEntry
	err := collection.FindOne(ctx, bson.M{"swift_code": swiftCode}).Decode(&result)
	if err != nil {
		log.Printf("FindBySwiftCode: Not found or error: %v\n", err)
		return nil, err
	}
	log.Printf("FindBySwiftCode: Found entry: %+v\n", result)
	return &result, nil
}

func FindBranchesByHqCode(ctx context.Context, collection *mongo.Collection, hqCode string) ([]models.SwiftEntry, error) {
	log.Printf("FindBranchesByHqCode: Searching for branches with hqCode = %s\n", hqCode)

	cursor, err := collection.Find(ctx, bson.M{"hqCode": hqCode})
	if err != nil {
		log.Printf("FindBranchesByHqCode: Query error: %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var branches []models.SwiftEntry
	if err := cursor.All(ctx, &branches); err != nil {
		log.Printf("FindBranchesByHqCode: Decoding error: %v\n", err)
		return nil, err
	}

	log.Printf("FindBranchesByHqCode: Found %d branches\n", len(branches))
	return branches, nil
}

func FindByCountryCode(ctx context.Context, collection *mongo.Collection, countryISO2 string) (*models.CountryShort, error) {
	log.Printf("FindByCountryCode: Searching for country with country_code = %s\n", countryISO2)

	// No separate country documents, so we fetch from the first available bank
	filter := bson.M{"country_code": countryISO2}
	projection := bson.M{
		"_id":          0,
		"country_code": 1,
		"country_name": 1,
	}
	opts := options.FindOne().SetProjection(projection)

	var result models.CountryShort
	err := collection.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		log.Printf("FindByCountryCode: Country not found: %v\n", err)
		return nil, err
	}

	log.Printf("FindByCountryCode: Found country: %+v\n", result)
	return &result, nil
}

func FindBanksByCountry(ctx context.Context, collection *mongo.Collection, countryISO2code string) ([]models.SwiftEntry, error) {
	log.Printf("FindBanksByCountry: Searching for banks with country_code = %s\n", countryISO2code)

	cursor, err := collection.Find(ctx, bson.M{"country_code": countryISO2code})
	if err != nil {
		log.Printf("FindBanksByCountry: Query error: %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var banks []models.SwiftEntry
	if err := cursor.All(ctx, &banks); err != nil {
		log.Printf("FindBanksByCountry: Decoding error: %v\n", err)
		return nil, err
	}

	log.Printf("FindBanksByCountry: Found %d banks\n", len(banks))
	return banks, nil
}
