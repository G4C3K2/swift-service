package main

import (
	"context"
	"log"

	"github.com/G4C3K2/swift-service/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Tworzymy opcje klienta z URI MongoDB
	clientOptions := options.Client().ApplyURI("mongodb+srv://kamilkoscielny2002:Qetuoadgjlxvn.1@cluster0.wpmnnvs.mongodb.net/")

	// Łączymy się z bazą
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Błąd połączenia z MongoDB:", err)
	}
	defer client.Disconnect(context.TODO())

	// Wybieramy bazę i kolekcję
	db := client.Database("swiftdb")
	collection := db.Collection("entries")

	// Ładujemy dane z pliku CSV
	utils.LoadData("Data.csv", collection)
}
