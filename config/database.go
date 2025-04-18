package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var DB *mongo.Database

func InitialDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file:", err)
		// Continue execution, as env vars might be set through other means
	}

	// Get MongoDB URI from environment variable
	dbURI := os.Getenv("MANGODB_URL")
	if dbURI == "" {
		log.Fatal("MANGODB_URL environment variable is not set")
	}

	fmt.Println("Connecting to MongoDB with URL:", dbURI)

	// Set up the MongoDB client options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(dbURI).
		SetServerAPIOptions(serverAPI).
		SetTimeout(30 * time.Second)

	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}

	// Ping the database with a timeout context
	pingCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
		log.Fatal("MongoDB Ping Error:", err)
	}

	// Set global DB variable
	dbName := os.Getenv("DATABASE")
	if dbName == "" {
		dbName = "bitslearn"
	}

	DB = client.Database(dbName)

	fmt.Println("Connected to MongoDB database:", dbName)
}

// Function to properly disconnect from MongoDB
func DisconnectDB(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		log.Fatal("Failed to disconnect from MongoDB:", err)
	}
	fmt.Println("Disconnected from MongoDB")
}
