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
		log.Fatal("MONGO_URI does not exist in environment variables")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("MongoDB does not respond:", err)
	}

	db := client.Database("swift") // Zmień na "swift"
	return db.Collection("entries")
}
