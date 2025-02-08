package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient holds the MongoDB client instance
type MongoClient struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// NewMongoClient initializes a new MongoDB client and connects to the database
func NewMongoClient(uri, dbName string) *MongoClient {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Verify the connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
	return &MongoClient{
		Client: client,
		DB:     client.Database(dbName),
	}
}

// Disconnect disconnects the MongoDB client
func (mc *MongoClient) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := mc.Client.Disconnect(ctx); err != nil {
		log.Fatalf("Failed to disconnect MongoDB client: %v", err)
	}

	log.Println("Disconnected from MongoDB")
}
