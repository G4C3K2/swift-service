package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConnection() *mongo.Collection {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("Brak MONGO_URI w zmiennych środowiskowych")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Błąd połączenia z MongoDB:", err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("MongoDB nie odpowiada:", err)
	}

	db := client.Database("swiftdb")
	return db.Collection("entries")
}
