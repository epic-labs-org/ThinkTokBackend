package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("ThinkTokDB").Collection("words")

	words := []interface{}{
		bson.M{"word": "apple", "translation": "سیب", "language": "English"},
		bson.M{"word": "house", "translation": "خانه", "language": "English"},
		bson.M{"word": "car", "translation": "ماشین", "language": "English"},
	}

	if _, err := collection.InsertMany(ctx, words); err != nil {
		log.Fatalf("Failed to seed words: %v", err)
	}
	log.Println("Global words seeded successfully")
}
