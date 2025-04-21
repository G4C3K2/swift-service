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
		log.Fatalf("Błąd tworzenia klienta MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Nie udało się połączyć z MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("swift").Collection("entries")

	var docs []interface{}
	for _, entry := range entries {
		docs = append(docs, entry)
	}

	_, err = collection.InsertMany(ctx, docs)
	if err != nil {
		log.Fatalf("Błąd przy zapisie do Mongo: %v", err)
	}
}

func FindBySwiftCode(ctx context.Context, collection *mongo.Collection, swiftCode string) (*models.SwiftEntry, error) {
	var result models.SwiftEntry
	err := collection.FindOne(ctx, bson.M{"swift_code": swiftCode}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func FindBranchesByHqCode(ctx context.Context, collection *mongo.Collection, hqCode string) ([]models.SwiftEntry, error) {
	cursor, err := collection.Find(ctx, bson.M{"hqCode": hqCode})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var branches []models.SwiftEntry
	if err := cursor.All(ctx, &branches); err != nil {
		return nil, err
	}

	return branches, nil
}
